package cli

import (
	"context"
	"os"

	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/Ein-Framework/Ein-Framework/pkg/log"
	"github.com/urfave/cli/v2"
)

const (
	CliName = "ein"
)

func CreateCli(cmds ...[]*cli.Command) *cli.App {
	var commands []*cli.Command

	for _, moreCommands := range cmds {
		commands = append(commands, moreCommands...)
	}

	app := &cli.App{
		Name: CliName,
		Before: func(ctx *cli.Context) error {

			debugEnable := os.Getenv("EIN_DEBUG")

			logger := log.New(debugEnable == "1")
			frameworkConfig, err := config.GetConfig()
			if err != nil {
				logger.Sugar().Panicln("[-] Error: loading configuration")
			}

			logger.Sugar().Infoln(frameworkConfig)

			ctx.Context = context.WithValue(ctx.Context, configKey, frameworkConfig)
			ctx.Context = context.WithValue(ctx.Context, loggerKey, logger)

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

