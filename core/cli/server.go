package cli

import (
	"github.com/Ein-Framework/Ein-Framework/core/api"
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/Ein-Framework/Ein-Framework/pkg/log"
	"github.com/urfave/cli/v2"
	"go.uber.org/zap"
)

func ServerCommands() []*cli.Command {

	startServerCommand := &cli.Command{
		Name:    "server",
		Aliases: []string{"svr"},
		Usage:   "Start HTTP server.",
		Action: func(ctx *cli.Context) error {
			frameworkConfig := ctx.Context.Value(configKey).(*config.Config)

			logger := ctx.Context.Value(loggerKey).(*zap.Logger)

			if frameworkConfig == nil {
				log.LogError("[-] error: config is not setup")
				return nil
			}
			if logger == nil {
				log.LogError("[-] error: config is not setup")
				return nil
			}

			api.New(*frameworkConfig,logger)

			return nil
		},
	}

	var commands []*cli.Command

	commands = append(commands, startServerCommand)

	return commands
}
