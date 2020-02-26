package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"strings"
	"trdeploy/flags"
)

func initAction(c *cli.Context, opts ...CommandOption) error {
	prefix := c.String(flags.Prefix)
	wp := c.String(flags.WorkProfile)
	initDir, initPath := getPaths(c.String(flags.ExecDir))

	terragruntConfigPath := initPath + "/" + terragruntConfigName
	terraformDirPath := initPath + "/" + ".terraform"
	terraformStatePath := fmt.Sprintf("%s/%s/%s_%s_%s.tfstate", wp, prefix, wp, prefix, initDir)

	tCfg := fmt.Sprintf(
		terragruntConfigTempl,
		c.String(flags.S3StateBacket),
		terraformStatePath,
		c.String(flags.Region),
		c.String(flags.DynamodbLockTable),
		c.String(flags.AuditProfile),
	)

	if err := os.RemoveAll(terraformDirPath); !os.IsNotExist(err) && err != nil {
		cli.Exit(fmt.Sprintf("delete .terraform error: %+v", err), 1)
		return fmt.Errorf("delete .terraform error: %+v", err)
	}

	if err := os.RemoveAll(terragruntConfigPath); !os.IsNotExist(err) && err != nil {
		cli.Exit(fmt.Sprintf("delete %s error: %+v", terragruntConfigName, err), 1)
		return fmt.Errorf("delete %s error: %+v", terragruntConfigName, err)
	}

	err := ioutil.WriteFile(terragruntConfigPath, []byte(tCfg), 0777)
	if err != nil {
		cli.Exit(fmt.Sprintf("creating terragrunt config error: %+v", err), 1)
		return fmt.Errorf("creating terragrunt config error: %+v", err)
	}

	return execute([]string{"init", "--terragrunt-config", terragruntConfigPath}, c, opts...)
}

func getPaths(path string) (string, string) {
	if path != "" {
		initPath := strings.Split(path, "/")
		return initPath[len(initPath)-1], currentPath() + "/" + path
	}

	return CurrentDir(), currentPath()
}
