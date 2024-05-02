package core

import (
	"context"
	"log"

	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/urfave/cli/v2"
)

type ConfigKey string

const (
	configKey ConfigKey = "config"
)

const (
	CliName = "ein"
)

func CreateApp(cmds ...[]*cli.Command) *cli.App {
	var commands []*cli.Command

	for _, moreCommands := range cmds {
		commands = append(commands, moreCommands...)
	}

	app := &cli.App{
		Name: CliName,
		Before: func(ctx *cli.Context) error {

			frameworkConfig, err := config.GetConfig()
			if err != nil {
				log.Panicf("Error: loading configuration")
			}

			ctx.Context = context.WithValue(ctx.Context, configKey, frameworkConfig)
			return nil

		},
		Commands: commands,
		Authors: []*cli.Author{
			{
				Name:  "Yassine Belkhadem",
				Email: "yassine.belkhadem@insat.rnu.tn",
			},
			{
				Name:  "Med Mongi Saidane",
				Email: "saidanemongi@gmail.com",
			},
			{
				Name:  "Salma Seddik",
				Email: "",
			},
			{
				Name:  "Hani Haded",
				Email: "",
			},
		},
		Usage:                "Automate your scans, sit & watch.",
		Copyright:            "(c) 2024 INSAT",
		EnableBashCompletion: true,
	}

	return app
}
