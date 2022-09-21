package app

import (
	"log"

	"git.dar.kz/forte-market/migrations/internal/app/config"
	"git.dar.kz/forte-market/migrations/internal/app/store"
	"git.dar.kz/forte-market/migrations/internal/migrate/presenter"
	"git.dar.kz/forte-market/migrations/internal/migrate/presenter/cmd"
	"github.com/golang-migrate/migrate/v4"
	"github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli"
)

func Run(args []string) {
	cfg, err := config.GetConfig()
	if err != nil {
		log.Fatal(err)
	}

	db, err := store.InitPostgresConnection(cfg.Postgres)
	if err != nil {
		log.Fatal("postgres: ", err)
	}
	defer func() {
		if err := db.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	driver, err := postgres.WithInstance(db, &postgres.Config{})
	if err != nil {
		log.Fatal("driver: ", err)
	}

	dbInstance, err := migrate.NewWithDatabaseInstance(cfg.MigrationsURL, cfg.Postgres.DB, driver)
	if err != nil {
		log.Fatal("migrate instance: ", err)
	}

	migrateService := presenter.New(dbInstance)
	app := &cli.App{
		Name:     "migrate cli",
		Commands: cmd.MakeHandler(migrateService),
	}

	if err = app.Run(args); err != nil {
		log.Fatal(err)
	}
}
