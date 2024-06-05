package cli

import (
	"fmt"

	"github.com/Ein-Framework/Ein-Framework/core/api"
	"github.com/Ein-Framework/Ein-Framework/core/domain"
	"github.com/Ein-Framework/Ein-Framework/core/services"
	taskmanager "github.com/Ein-Framework/Ein-Framework/core/task_manager"
	"github.com/Ein-Framework/Ein-Framework/core/templating"
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

			db, err := domain.NewDatabase(frameworkConfig.Database)

			if err != nil {
				log.Fatal(err.Error())
			}

			fmt.Println(db)

			// Initialize services
			coreServices := services.InitServices(db, logger, frameworkConfig)

			components := InitComponents(coreServices, frameworkConfig, logger)
			api.New(coreServices, components, frameworkConfig, logger)
			return nil
		},
	}

	var commands []*cli.Command

	commands = append(commands, startServerCommand)
	return commands
}

func InitComponents(coreServices *services.Services, frameworkConfig *config.Config, logger *zap.Logger) *api.AppComponents {
	templatingManager := templating.New(frameworkConfig, coreServices, logger)

	components := &api.AppComponents{
		TemplatingManager: templatingManager,
		TaskManager:       taskmanager.New(templatingManager, coreServices, logger),
	}
	return components
}
