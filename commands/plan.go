package commands

import (
	"github.com/urfave/cli/v2"
)

func plan(c *cli.Context, opts ...CommandOption) error {
	command := []string{"plan"}

	opts = append(opts, Cmd(c))

	return execute(command, c, opts...)
}
