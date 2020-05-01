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

	if c.IsSet(flags.AdditionalArgs) {
		command = append(command, c.String(flags.AdditionalArgs))
	}

	return execute(command, c, opts...)
}
