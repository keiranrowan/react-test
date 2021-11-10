package command

import (
	"github.com/urfave/cli/v2"
	"github.com/keiranrowan/react-test/internal/server"
)

var ServerCommand = cli.Command{
	Name: "server",
	Aliases: []string{"s"},
	Usage: "Server-side web server commands",
	Subcommands: []*cli.Command{
		{
            Name:  "start",
            Usage: "Start the server-side web server",
            Action: server.Init,
		},
	},
}
