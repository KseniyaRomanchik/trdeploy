package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"os/exec"
)

func InitAction (c *cli.Context) error {
	err := os.RemoveAll(".terraform")
	if !os.IsNotExist(err) && err != nil {
		cli.Exit("delete error", 86)
		return fmt.Errorf("delete error1: %+v", err)
	}

	err = os.RemoveAll(".terraform.tfstate")
	if !os.IsNotExist(err) && err != nil {
		cli.Exit("delete error", 86)
		return fmt.Errorf("delete error2: %+v", err)
	}

	err = os.RemoveAll(terragruntConfigName)
	if !os.IsNotExist(err) && err != nil {
		cli.Exit("delete error", 86)
		return fmt.Errorf("delete error3: %+v", err)
	}

	//prepare_terragrunt_config

	prefix := c.String("prefix")
	wp := c.String("work_profile")

	tCfg := fmt.Sprintf(
		terragruntConfigTempl,
		c.String("s3_state_backet"),
		fmt.Sprintf("%s/%s/%s_%s_%s.tfstate", wp, prefix, wp, prefix, CurrentDir()),
		c.String("region"),
		c.String("dynamodb_lock_table"),
		c.String("audit_profile"),
	)

	err = ioutil.WriteFile(terragruntConfigName, []byte(tCfg), 0777)
	if err != nil {
		cli.Exit("creating terragrunt config error", 86)
		return fmt.Errorf("creating terragrunt config error: %+v", err)
	}

	cmdInit := exec.Command("terragrunt", "init",  "--terragrunt-config", terragruntConfigName)

	outInit, err := cmdInit.CombinedOutput()
	if err != nil {
		cli.Exit("terragrunt error", 86)
		return fmt.Errorf("terragrunt error: %+v \n %s", err, outInit)
	}

	return nil
}