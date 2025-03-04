// Package bindings is responsible for providing a bridge between the code that
// is generated from the OpenAPI specification that comprises the "Transport"
// layer and the code that implements the product, or the "Service" layer.
//
// This package is structured as a set of structs which are compose together in
// the struct called `Bindings` which implements an interface generated by the
// oapi-codegen tool which describes all the endpoints.
//
// Aside from `bindings.go`, pretty much every Go file is named after a REST
// collection from the OpenAPI specification.
//
// ## Adding Routes
//
// To add a new route, you first modify the OpenAPI specification YAML document
// and then run the `generate` task which generates all the handler declarations
// and types necessary to implement the binding.
//
// Next, you create a file in this package named after the collection. So if you
// added `/v1/things` you'd create `things.go` and inside that file, a struct
// named `Things` and a constructor named `NewThings`. This pattern may not
// always apply for certain cases but it's generally best to try to follow.
//
// You then add your struct to the `Bindings` composed struct and provide the
// implementation of your struct to the DI system using `bindingsProviders`.
//
// ## Changing Routes
//
// Updating a route is as simple as just modifying the OpenAPI specification
// and making the necessary changes to the bindings to get the code compiling.
package bindings

import (
	"context"
	"net/http"
	"strings"

	"github.com/Southclaws/fault"
	"github.com/Southclaws/fault/fmsg"
	oapi_middleware "github.com/deepmap/oapi-codegen/pkg/middleware"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/samber/lo"
	"go.uber.org/fx"
	"go.uber.org/zap"

	"github.com/Southclaws/storyden/app/transports/openapi/glue"
	"github.com/Southclaws/storyden/internal/config"
	"github.com/Southclaws/storyden/internal/openapi"
)

// Bindings is a DI parameter struct that is used to compose together all of the
// individual service bindings in this package. When the provider below depends
// on this type, it provides all these composed bindings to the DI system so the
// invoke call can mount them onto the router using the `StrictServerInterface`.
//
// The reason this is done this way is so we split code up based on OpenAPI
// REST collections instead of bundling everything into one huge struct with
// loads of dependencies. This is just how the oapi-codegen tool works, by
// generating one big interface which the bindings layer must satisfy.
type Bindings struct {
	fx.In
	Version
	Spec
	Authentication
	WebAuthn
	Accounts
	Profiles
	Threads
	Posts
}

// bindingsProviders provides to the application the necessary implementations
// that compose the `Bindings` parameter struct which implements the OpenAPI
// server interface. When you add a new collection, add it to Bindings and here.
func bindingsProviders() fx.Option {
	return fx.Provide(
		NewVersion,
		NewSpec,
		NewAuthentication,
		NewWebAuthn,
		NewAccounts,
		NewProfiles,
		NewThreads,
		NewPosts,
	)
}

// bindings provides to the application the above struct which binds the service
// layer to the transport layer. This uses `Bindings` as an fx parameter struct.
//
// ## WHY AM I GETTING AN ERROR HERE?
//
// When you edit `openapi.yaml` and re-run the code generation task, this will
// most likely change the declaration of `StrictServerInterface` inside the
// generated package `openapi`.
//
// The error you will see is most likely something along the lines of:
//
//	*Bindings does not implement openapi.StrictServerInterface
//
// and the underlying problem is either missing methods or methods that have
// changed signature due to changes to the parameters or request or response.
//
// This API follows RESTful design so a collection in the API specification
// (such as `/v1/accounts`) will map to a file, struct and constructor here
// (such as `accounts.go`, `Accounts` and `NewAccounts`) and everything is glued
// together in this file.
func bindings(s Bindings) openapi.StrictServerInterface {
	return &s
}

// mounts the OpenAPI routes and middleware onto the /api path. Everything that
// is outside of the `/api` path is considered part of the proxied frontend app.
// Note: routes are mounted with the `OnStart` hook so that middleware is first.
func mount(lc fx.Lifecycle, l *zap.Logger, mux *http.ServeMux, router *echo.Echo, si openapi.StrictServerInterface) {
	lc.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			openapi.RegisterHandlersWithBaseURL(router, openapi.NewStrictHandler(si, nil), "/api")

			l.Info("mounted OpenAPI to service bindings",
				zap.Strings("routes", lo.Map(router.Routes(), func(r *echo.Route, _ int) string {
					return r.Path
				})),
			)

			// mount onto / because this router already only cares about /api
			mux.Handle("/", router)

			return nil
		},
	})
}

func addMiddleware(cfg config.Config, l *zap.Logger, router *echo.Echo, auth Authentication) error {
	spec, err := openapi.GetSwagger()
	if err != nil {
		return fault.Wrap(err, fmsg.With("failed to get openapi specification"))
	}

	router.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			r := c.Request()

			l.Info(
				"request",
				zap.String("method", r.Method),
				zap.String("path", r.URL.Path),
				zap.Any("query", r.URL.Query()),
				zap.Int64("body", r.ContentLength),
			)

			return next(c)
		}
	})

	// router.Use(echo.WrapMiddleware(func(h http.Handler) http.Handler {
	// 	proxy := httputil.NewSingleHostReverseProxy(utils.Must(url.Parse("http://localhost:3000")))

	// 	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
	// 		if !strings.HasPrefix(r.URL.Path, "/api") {
	// 			fmt.Println("PROXY 2 FRONT END PLS")

	// 			proxy.ServeHTTP(w, r)

	// 			return
	// 		}

	// 		h.ServeHTTP(w, r)
	// 	})
	// }))

	router.Use(auth.middleware)
	router.Use(oapi_middleware.OapiRequestValidatorWithOptions(spec, &oapi_middleware.Options{
		Skipper: openApiSkipper,
		Options: openapi3filter.Options{
			IncludeResponseStatus: true,
			AuthenticationFunc:    auth.validator,
		},
		// Handles validation errors that occur BEFORE the handler is called.
		ErrorHandler: glue.ValidatorErrorHandler(),
	}))

	return nil
}

func openApiSkipper(c echo.Context) bool {
	return !strings.HasPrefix(c.Path(), "/api")
}

func newRouter(l *zap.Logger, cfg config.Config) *echo.Echo {
	router := echo.New()

	origins := []string{
		"http://localhost:3000", // Local development
		"http://localhost:8001", // Swagger UI
		cfg.PublicWebAddress,    // Live public website
	}

	router.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins:     origins,
		AllowMethods:     []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Accept", "Authorization", "Content-Type", "Content-Length", "X-CSRF-Token"},
		ExposeHeaders:    []string{"Link", "Content-Length", "X-Ratelimit-Limit", "X-Ratelimit-Reset"},
		AllowCredentials: true,
		MaxAge:           300,
	}))

	router.Use(glue.ParameterContext)
	router.HTTPErrorHandler = glue.HTTPErrorHandler(l)

	// Router must add all middleware before mounting routes. To add middleware,
	// simply depend on the router in a provider or invoker and do `router.Use`.
	// To mount routes use the lifecycle `OnStart` hook and mount them normally.

	l.Info("created router", zap.Strings("origins", origins))

	return router
}

func Build() fx.Option {
	return fx.Options(
		// Provide the bindings struct which implements the generated OpenAPI
		// interface by composing together all of the service bindings into a
		// single struct.
		fx.Provide(bindings),

		// Provide the Echo router.
		fx.Provide(newRouter),

		// Add the middleware bindings.
		fx.Invoke(addMiddleware),

		// Mount the bound OpenAPI routes onto the router.
		fx.Invoke(mount),

		// Provide all service layer bindings to the DI system.
		bindingsProviders(),
	)
}
