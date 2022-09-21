package cmd

import (
	"context"
	"fmt"
	"time"

	"git.dar.kz/forte-market/migrations/internal/migrate/presenter"
	"github.com/golang-migrate/migrate/v4"
	"github.com/urfave/cli"
)

func makeMigrateUpEndpoint(service presenter.Service) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		ctx, ctxCancel := context.WithTimeout(context.Background(), time.Minute*10)
		defer ctxCancel()

		err := service.MigrateUp(ctx)
		if err != nil {
			if err == migrate.ErrNoChange {
				fmt.Println(err.Error())
				return nil
			}
			return err
		}
		return nil
	}
}

func makeMigrateVersionEndpoint(service presenter.Service) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		version, err := service.MigrateVersion()
		if err != nil {
			return err
		}
		fmt.Println(version)
		return nil
	}
}

func makeMigrateDropEndpoint(service presenter.Service) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		err := service.MigrateDrop()
		if err != nil {
			return err
		}
		return nil
	}
}

func makeGoToVersionEndpoint(service presenter.Service) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		ctx, ctxCancel := context.WithTimeout(context.Background(), time.Minute*10)
		defer ctxCancel()

		version := decodeGoToVersionRequest(c)
		if version == 0 {
			return fmt.Errorf("")
		}

		err := service.GoToVersion(ctx, version)
		if err != nil {
			return err
		}
		return nil
	}
}

func decodeGoToVersionRequest(c *cli.Context) uint {
	return c.Uint("version")
}
