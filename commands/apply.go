package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
	"trdeploy/flags"
)

func apply(c *cli.Context) error {
	planFile := c.String(flags.PlanFile)
	prefix := c.String(flags.Prefix)
	additionalArgs := c.String(flags.AdditionalArgs)
	ap := c.String(flags.AuditProfile)
	wp := c.String(flags.WorkProfile)
	gvp := c.String(flags.GlobalVarPath)
	mtfv := c.String(flags.ModuleTfvars)

	if c.IsSet(flags.PlanFile) {
		cmd := exec.Command("terragrunt", "apply", "--terragrunt-config", planFile)

		out, err := cmd.CombinedOutput()
		if err != nil {
			cli.Exit("terragrunt error", 1)
			return fmt.Errorf("terragrunt error: %+v \n %s", err, out)
		}

		return nil
	}

	cmdApply := exec.Command(
		"terragrunt", "apply",
		"-var-file", fmt.Sprintf("%s/common.tfvars", gvp),
		"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
		"-var-file", mtfv,
		"-var", fmt.Sprintf("prefix=%s", prefix),
		"-var", fmt.Sprintf("aws_audit=%s", ap),
		additionalArgs,
	)

	outApply, err := cmdApply.CombinedOutput()
	if err != nil {
		cli.Exit("terragrunt apply error", 1)
		return fmt.Errorf("\nterragrunt apply error: %+v \n%s", err, outApply)
	}

	return nil
}
