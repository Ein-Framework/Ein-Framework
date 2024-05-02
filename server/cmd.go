package server

import (
	"github.com/Ein-Framework/Ein-Framework/pkg/config"
	"github.com/urfave/cli/v2"
)

func GetCommands() []*cli.Command {

	startServerCommand := &cli.Command{
		Name:    "server",
		Aliases: []string{"svr"},
		Usage:   "Start HTTP server.",
		Action: func(cCtx *cli.Context) error {
			frameworkConfig := cCtx.Context.Value("config").(*config.Config)

			if frameworkConfig == nil {
				return nil
			}

			// var serverManager IServerManager = &Manager{}
			go func() {
				<-make(chan int)
				// err := serverManager.NewgRPCServer(frameworkConfig)
				// if err != nil {
				// 	log.Println("Error launching GRPC server")
				// 	return
				// }
			}()

			return nil
		},
	}

	var commands []*cli.Command

	commands = append(commands, startServerCommand)

	return commands
}
