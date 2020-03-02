package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strconv"
	"trdeploy/flags"
)

func destroy(c *cli.Context, opts ...CommandOption) error {
	prefix := c.String(flags.Prefix)
	ap := c.String(flags.AuditProfile)
	wp := c.String(flags.WorkProfile)
	gvp := c.String(flags.GlobalVarPath)
	mtfv := c.String(flags.ModuleTfvars)
	p := c.Int(flags.Parallelism)

	commands := []string{
		"destroy",
		"-var-file", fmt.Sprintf("%s/common.tfvars", gvp),
		"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
		"-var-file", mtfv,
		"-var", fmt.Sprintf("prefix=%s", prefix),
		"-var", fmt.Sprintf("aws_audit=%s", ap),
	}

	if !c.IsSet(flags.Parallelism) {
		p = defaultDestroyOperations
	}

	commands = append(commands, "-parallelism", strconv.Itoa(p))

	return execute(commands, c, opts...)
}
