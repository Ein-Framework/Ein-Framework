package core

import (
	einCli "github.com/Ein-Framework/Ein-Framework/core/cli"

	"github.com/urfave/cli/v2"
)

func Setup() *cli.App {
	app := einCli.CreateCli(einCli.ConfigCommands(), einCli.ServerCommands())
	return app
}
