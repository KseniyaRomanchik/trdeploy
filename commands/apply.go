package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"strconv"
	"trdeploy/flags"
)

func apply(c *cli.Context, opts ...CommandOption) error {
	planFile := c.String(flags.PlanFile)
	prefix := c.String(flags.Prefix)
	ap := c.String(flags.AuditProfile)
	wp := c.String(flags.WorkProfile)
	gvp := c.String(flags.GlobalVarPath)
	mtfv := c.String(flags.ModuleTfvars)
	p := c.Int(flags.Parallelism)

	if c.IsSet(flags.PlanFile) {
		return execute([]string{"apply", planFile}, c, opts...)
	}

	commands := []string{
		"apply",
		"-var-file", fmt.Sprintf("%s/common.tfvars", gvp),
		"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
		"-var-file", mtfv,
		"-var", fmt.Sprintf("prefix=%s", prefix),
		"-var", fmt.Sprintf("aws_audit=%s", ap),
	}

	if c.IsSet(flags.Parallelism) {
		commands = append(commands, "-parallelism", strconv.Itoa(p))
	}

	return execute(commands, c, opts...)
}
