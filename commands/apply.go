package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
)

func Apply(c *cli.Context) error {
	planFile := c.String("plan_file")
	prefix := c.String("prefix")
	additionalArgs := c.String("additional_args")
	ap := c.String("audit_profile")
	wp := c.String("work_profile")
	gvp := c.String("global_var_path")
	mtfv := c.String("module_tfvars")

	if c.IsSet("plan_file") {
		cmd := exec.Command("terragrunt", "apply",  "--terragrunt-config", planFile)

		out, err := cmd.CombinedOutput()
		if err != nil {
			cli.Exit("terragrunt error", 86)
			return fmt.Errorf("terragrunt error: %+v \n %s", err, out)
		}

		return nil
	}

	//cmd := exec.Command(
	//	"terragrunt", "refresh",
	//	"-var-file", fmt.Sprintf( "%s/common.tfvars", gvp),
	//	"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
	//	"-var-file", mtfv,
	//	"-var", fmt.Sprintf("prefix=%s", prefix),
	//	"-var", fmt.Sprintf("aws_audit=%s", ap),
	//	additionalArgs,
	//)

	//outRefresh, err := cmd.CombinedOutput()
	//if err != nil {
	//	cli.Exit("terragrunt error", 86)
	//	return fmt.Errorf("terragrunt error: %+v \n %s", err, outRefresh)
	//}

	cmdApply := exec.Command(
		"terragrunt", "apply",
		"-var-file", fmt.Sprintf( "%s/common.tfvars", gvp),
		"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
		"-var-file", mtfv,
		"-var", fmt.Sprintf("prefix=%s", prefix),
		"-var", fmt.Sprintf("aws_audit=%s", ap),
		additionalArgs,
	)

	outApply, err := cmdApply.CombinedOutput()
	if err != nil {
		cli.Exit("terragrunt error", 86)
		return fmt.Errorf("terragrunt error: %+v \n %s", err, outApply)
	}

	return nil
}
