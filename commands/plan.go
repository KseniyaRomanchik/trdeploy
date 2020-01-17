package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os/exec"
	"trdeploy/flags"
)

func plan(c *cli.Context) error {
	prefix := c.String(flags.Prefix)
	additionalArgs := c.String(flags.AdditionalArgs)
	ap := c.String(flags.AuditProfile)
	wp := c.String(flags.WorkProfile)
	gvp := c.String(flags.GlobalVarPath)
	mtfv := c.String(flags.ModuleTfvars)
	outPlanLog := c.String(flags.OutPlanLog)

	cmdPlan := exec.Command(
		"terragrunt", "plan",
		"-var-file", fmt.Sprintf("%s/common.tfvars", gvp),
		"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
		"-var-file", mtfv,
		"-var", fmt.Sprintf("prefix=%s", prefix),
		"-var", fmt.Sprintf("aws_audit=%s", ap),
		additionalArgs,
	)

	outPlan, err := cmdPlan.CombinedOutput()
	if err != nil {
		cli.Exit("terragrunt plan error", 1)
		return fmt.Errorf("terragrunt error: %+v \n %s", err, outPlan)
	}

	if c.IsSet(flags.OutPlanLog) {
		err = ioutil.WriteFile(outPlanLog, []byte(outPlan), 0777)
		if err != nil {
			cli.Exit("creating out plan log file error", 1)
			return fmt.Errorf("creating out plan log file error: %+v", err)
		}
	}

	return nil
}
