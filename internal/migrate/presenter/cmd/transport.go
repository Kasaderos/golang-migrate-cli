package cmd

import (
	"git.dar.kz/forte-market/migrations/internal/migrate/presenter"
	"github.com/urfave/cli"
)

func MakeHandler(service presenter.Service) []cli.Command {
	return []cli.Command{
		{
			Name:   "up",
			Usage:  "up to the last migration version",
			Action: makeMigrateUpEndpoint(service),
		},
		{
			Name:    "goto",
			Aliases: []string{"gt"},
			Usage:   "migrates either up or down to the specified version",
			Action:  makeGoToVersionEndpoint(service),
			Flags: []cli.Flag{cli.UintFlag{
				Name:     "version",
				Required: true,
			}},
		},
		{
			Name:    "version",
			Aliases: []string{"v"},
			Usage:   "Print current migration version",
			Action:  makeMigrateVersionEndpoint(service),
		},
		{
			Name:   "drop",
			Usage:  "drops EVERTHING inside database",
			Action: makeMigrateDropEndpoint(service),
		},
	}
}
