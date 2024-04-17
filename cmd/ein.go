package main

import (
	"log"
	"os"

	"github.com/Ein-Framework/Ein-Framework/core"
	"github.com/Ein-Framework/Ein-Framework/core/config"
)

func main() {
	configCommands := config.GetCommands()

	app := core.CreateApp(configCommands)
	err := app.Run(os.Args)
	if err != nil {
		log.Panicln("Error occurred while launching the framework", err)
	}
}
