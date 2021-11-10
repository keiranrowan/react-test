package command

import (
	"github.com/urfave/cli/v2"
	"github.com/keiranrowan/react-test/internal/daemon"
)

var DaemonCommand = cli.Command{
	Name: "daemon",
	Aliases: []string{"d"},
	Usage: "Application daemon commands",
	Subcommands: []*cli.Command{
		{
            Name:  "start",
            Usage: "Start the application daemon",
            Action: daemon.Init,
		},
	},
}
