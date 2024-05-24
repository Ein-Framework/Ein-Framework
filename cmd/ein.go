package main

import (
	"log"
	"os"

	"github.com/Ein-Framework/Ein-Framework/core"
)

func main() {

	app := core.Setup()
	err := app.Run(os.Args)
	if err != nil {
		log.Panicln("Error occurred while launching the framework", err)
	}
}
