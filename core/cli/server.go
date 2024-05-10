package cli

import (
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/Ein-Framework/Ein-Framework/pkg/log"
	"github.com/urfave/cli/v2"
)

func ServerCommands() []*cli.Command {

	startServerCommand := &cli.Command{
		Name:    "server",
		Aliases: []string{"svr"},
		Usage:   "Start HTTP server.",
		Action: func(ctx *cli.Context) error {
			frameworkConfig := ctx.Context.Value(configKey).(*config.Config)

			if frameworkConfig == nil {
				return nil
			}

			log.LogError("Server not implemented Yet")

			return nil
		},
	}

	var commands []*cli.Command

	commands = append(commands, startServerCommand)

	return commands
}
