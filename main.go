package main

import (
	"os"
	"log"

	"github.com/urfave/cli/v2"
	"github.com/keiranrowan/react-test/internal/command"
)

func main() {
	app := &cli.App{
		Name: "react-test",
		Version: "0.0.1",
		Flags: []cli.Flag{

		},
		Commands: []*cli.Command {
			&command.ClientCommand,
			&command.ServerCommand,
			&command.DaemonCommand,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
