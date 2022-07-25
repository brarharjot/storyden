package http

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"go.uber.org/zap"

	"github.com/Southclaws/storyden/backend/internal/config"
)

func newRouter(l *zap.Logger, cfg config.Config) chi.Router {
	router := chi.NewRouter()

	origins := []string{
		"http://localhost:3000", // Local development
		cfg.PublicWebAddress,    // Live public website
	}

	router.Use(
		cors.Handler(cors.Options{
			AllowedOrigins:   origins,
			AllowedMethods:   []string{"GET", "POST", "PUT", "PATCH", "DELETE", "OPTIONS"},
			AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "Content-Length", "X-CSRF-Token"},
			ExposedHeaders:   []string{"Link", "Content-Length", "X-Ratelimit-Limit", "X-Ratelimit-Reset"},
			AllowCredentials: true,
			MaxAge:           300,
		}),
	)

	// Router must add all middleware before mounting routes. To add middleware,
	// simply depend on the router in a provider or invoker and do `router.Use`.
	// To mount routes use the lifecycle `OnStart` hook and mount them normally.

	l.Info("created router", zap.Strings("origins", origins))

	return router
}
