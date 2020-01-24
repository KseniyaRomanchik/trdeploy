package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"trdeploy/flags"
)

func apply(c *cli.Context) error {
	planFile := c.String(flags.PlanFile)
	prefix := c.String(flags.Prefix)
	ap := c.String(flags.AuditProfile)
	wp := c.String(flags.WorkProfile)
	gvp := c.String(flags.GlobalVarPath)
	mtfv := c.String(flags.ModuleTfvars)

	if c.IsSet(flags.PlanFile) {
		return execute([]string{"apply", planFile}, c)
	}

	return execute([]string{
		"apply",
		"-var-file", fmt.Sprintf("%s/common.tfvars", gvp),
		"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
		"-var-file", mtfv,
		"-var", fmt.Sprintf("prefix=%s", prefix),
		"-var", fmt.Sprintf("aws_audit=%s", ap),
	}, c)
}
