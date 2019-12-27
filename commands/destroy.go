package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
)

func Destroy(c *cli.Context) error {
	prefix := c.String("prefix")
	additionalArgs := c.String("additional_args")
	ap := c.String("audit_profile")
	wp := c.String("work_profile")
	gvp := c.String("global_var_path")
	mtfv := c.String("module_tfvars")

	cmdDestroy := exec.Command(
		"terragrunt", "destroy",
		"-var-file", fmt.Sprintf( "%s/common.tfvars", gvp),
		"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
		"-var-file", mtfv,
		"-var", fmt.Sprintf("prefix=%s", prefix),
		"-var", fmt.Sprintf("aws_audit=%s", ap),
		"-parallelism", "1",
		additionalArgs,
	)

	outDestroy, err := cmdDestroy.CombinedOutput()
	if err != nil {
		cli.Exit("terragrunt error", 86)
		return fmt.Errorf("terragrunt error: %+v \n %s", err, outDestroy)
	}

	return nil
}
