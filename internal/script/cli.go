package script

import (
	"context"
	"fmt"
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"
	"go.uber.org/fx"

	"github.com/Southclaws/storyden/app/resources"
	"github.com/Southclaws/storyden/app/services"
	"github.com/Southclaws/storyden/internal/config"
	"github.com/Southclaws/storyden/internal/infrastructure"
)

// Run is a quick helper for writing scripts that use services.
func Run(opts ...fx.Option) {
	ctx, cf := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cf()

	all := append(opts, []fx.Option{
		fx.NopLogger,
		config.Build(),
		infrastructure.Build(),
		services.Build(),
		resources.Build(),
	}...)

	app := fx.New(all...)

	if err := app.Start(ctx); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
