package db

import (
	"context"
	"database/sql"

	"entgo.io/ent/dialect"             // nolint:gci
	entsql "entgo.io/ent/dialect/sql"  // nolint:gci
	"entgo.io/ent/dialect/sql/schema"  // nolint:gci
	"github.com/Southclaws/fault"      // nolint:gci
	"github.com/Southclaws/fault/fctx" // nolint:gci
	"github.com/Southclaws/fault/fmsg" // nolint:gci
	_ "github.com/jackc/pgx/v4/stdlib" // nolint:gci
	"go.uber.org/fx"                   // nolint:gci

	"github.com/Southclaws/storyden/internal/config"
	"github.com/Southclaws/storyden/internal/ent"
)

func Build() fx.Option {
	return fx.Options(
		// provide the underlying *sql.DB to the system
		fx.Provide(newSQL),

		// construct a new ent client using the *sql.DB provided above
		fx.Provide(newEntClient),
	)
}

func newSQL(cfg config.Config) (*sql.DB, error) {
	driver, err := sql.Open("pgx", cfg.DatabaseURL)
	if err != nil {
		return nil, fault.Wrap(err, fmsg.With("failed to connect to database"))
	}

	return driver, nil
}

func newEntClient(lc fx.Lifecycle, db *sql.DB) (*ent.Client, error) {
	wctx, cancel := context.WithCancel(context.Background())

	client, err := connect(wctx, db)
	if err != nil {
		cancel()
		return nil, err
	}

	lc.Append(fx.Hook{
		OnStop: func(ctx context.Context) error {
			defer cancel()

			err := client.Close()
			if err != nil {
				return fault.Wrap(err, fctx.With(ctx))
			}

			return nil
		},
	})

	return client, nil
}

func connect(ctx context.Context, driver *sql.DB) (*ent.Client, error) {
	client := ent.NewClient(ent.Driver(entsql.OpenDB(dialect.Postgres, driver)))

	opts := []schema.MigrateOption{
		schema.WithAtlas(true),
	}

	// We don't do versioned migrations currently.
	// opts = append(opts, schema.WithDropColumn(true))
	// opts = append(opts, schema.WithDropIndex(true))

	// Run only additive migrations
	if err := client.Schema.Create(ctx, opts...); err != nil {
		return nil, fault.Wrap(err)
	}

	return client, nil
}
