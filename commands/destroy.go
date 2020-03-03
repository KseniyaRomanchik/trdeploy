package commands

import (
	"github.com/urfave/cli/v2"
	"trdeploy/flags"
)

func destroy(c *cli.Context, opts ...CommandOption) error {
	p := defaultDestroyOperations
	if c.IsSet(flags.Parallelism) {
		p = c.Int(flags.Parallelism)
	}

	command := []string{"destroy"}

	opts = append(opts, Cmd(c), Parallelism(p))

	return execute(command, c, opts...)
}
