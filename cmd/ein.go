package main

import (
	"log"
	"os"

	"github.com/Ein-Framework/Ein-Framework/core"
	"github.com/Ein-Framework/Ein-Framework/core/config"
	"github.com/Ein-Framework/Ein-Framework/server"
)

func main() {
	configCommands := config.GetCommands()
	serverCommands := server.GetCommands()

	app := core.CreateApp(configCommands, serverCommands)
	err := app.Run(os.Args)
	if err != nil {
		log.Panicln("Error occurred while launching the framework", err)
	}
}
