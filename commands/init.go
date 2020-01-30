package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"trdeploy/flags"
)

func initAction(c *cli.Context, opts ...CommandOption) error {
	err := os.RemoveAll(".terraform")
	if !os.IsNotExist(err) && err != nil {
		cli.Exit("delete .terraform error", 1)
		return fmt.Errorf("delete .terraform error: %+v", err)
	}

	err = os.RemoveAll(".terraform.tfstate")
	if !os.IsNotExist(err) && err != nil {
		cli.Exit("delete .terraform.tfstate error", 1)
		return fmt.Errorf("delete .terraform.tfstate error: %+v", err)
	}

	err = os.RemoveAll(terragruntConfigName)
	if !os.IsNotExist(err) && err != nil {
		cli.Exit(fmt.Sprintf("delete %s error", terragruntConfigName), 1)
		return fmt.Errorf("delete %s error: %+v", terragruntConfigName, err)
	}

	//prepare_terragrunt_config

	prefix := c.String(flags.Prefix)
	wp := c.String(flags.WorkProfile)

	tCfg := fmt.Sprintf(
		terragruntConfigTempl,
		c.String(flags.S3StateBacket),
		fmt.Sprintf("%s/%s/%s_%s_%s.tfstate", wp, prefix, wp, prefix, CurrentDir()),
		c.String(flags.Region),
		c.String(flags.DynamodbLockTable),
		c.String(flags.AuditProfile),
	)

	err = ioutil.WriteFile(terragruntConfigName, []byte(tCfg), 0777)
	if err != nil {
		cli.Exit("creating terragrunt config error", 1)
		return fmt.Errorf("creating terragrunt config error: %+v", err)
	}

	return execute([]string{
		"init",
	}, c, opts...)
}
