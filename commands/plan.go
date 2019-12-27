package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os/exec"
)

func Plan(c *cli.Context) error {
	prefix := c.String("prefix")
	additionalArgs := c.String("additional_args")
	ap := c.String("audit_profile")
	wp := c.String("work_profile")
	gvp := c.String("global_var_path")
	mtfv := c.String("module_tfvars")
	outPlanLog := c.String("out_plan_log")

	cmdPlan := exec.Command(
		"terragrunt", "plan",
		"-var-file", fmt.Sprintf( "%s/common.tfvars", gvp),
		"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
		"-var-file", mtfv,
		"-var", fmt.Sprintf("prefix=%s", prefix),
		"-var", fmt.Sprintf("aws_audit=%s", ap),
		additionalArgs,
	)

	outPlan, err := cmdPlan.CombinedOutput()
	if err != nil {
		cli.Exit("terragrunt error", 86)
		return fmt.Errorf("terragrunt error: %+v \n %s", err, outPlan)
	}

	if c.IsSet("out_plan_log") {
		err = ioutil.WriteFile(outPlanLog, []byte(outPlan), 0777)
		if err != nil {
			cli.Exit("creating out plan log file error", 86)
			return fmt.Errorf("creating out plan log file error: %+v", err)
		}
	}

	return nil
}
