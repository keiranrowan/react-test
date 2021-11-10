package command

import (
	"github.com/urfave/cli/v2"
	"github.com/keiranrowan/react-test/internal/client"
)

var ClientCommand = cli.Command{
	Name: "client",
	Aliases: []string{"c"},
	Usage: "Client-side web server commands",
	Subcommands: []*cli.Command{
		{
            Name:  "start",
            Usage: "Start the client-side web server",
            Action: client.Init,
		},
	},
}
