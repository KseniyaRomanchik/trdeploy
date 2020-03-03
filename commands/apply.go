package commands

import (
	"github.com/urfave/cli/v2"
	"trdeploy/flags"
)

func apply(c *cli.Context, opts ...CommandOption) error {
	planFile := c.String(flags.PlanFile)
	p := c.Int(flags.Parallelism)

	command := []string{"apply"}

	if c.IsSet(flags.Parallelism) {
		opts = append(opts, Parallelism(p))
	}

	if c.IsSet(flags.PlanFile) {
		return execute(append(command, planFile), c, opts...)
	}

	opts = append(opts, Cmd(c))

	return execute(command, c, opts...)
}
