package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
	"trdeploy/flags"
)

func destroy(c *cli.Context) error {
	prefix := c.String(flags.Prefix)
	additionalArgs := c.String(flags.AdditionalArgs)
	ap := c.String(flags.AuditProfile)
	wp := c.String(flags.WorkProfile)
	gvp := c.String(flags.GlobalVarPath)
	mtfv := c.String(flags.ModuleTfvars)

	cmdDestroy := exec.Command(
		"terragrunt", "destroy",
		"-var-file", fmt.Sprintf("%s/common.tfvars", gvp),
		"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
		"-var-file", mtfv,
		"-var", fmt.Sprintf("prefix=%s", prefix),
		"-var", fmt.Sprintf("aws_audit=%s", ap),
		"-parallelism", "1",
		additionalArgs,
	)

	outDestroy, err := cmdDestroy.CombinedOutput()
	if err != nil {
		cli.Exit("terragrunt destroy error", 1)
		return fmt.Errorf("terragrunt error: %+v \n %s", err, outDestroy)
	}

	return nil
}
