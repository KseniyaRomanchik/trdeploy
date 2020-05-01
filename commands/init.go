package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
	"trdeploy/flags"
)

func initAction(c *cli.Context, opts ...CommandOption) error {
	prefix := c.String(flags.Prefix)
	wp := c.String(flags.WorkProfile)

	cmd := Command{Cmd: &exec.Cmd{}}
	for _, opt := range opts {
		cmd = opt(cmd)
	}

	initDir, initPath := getPaths(cmd.Dir)

	terragruntConfigPath := initPath + "/" + TerragruntConfigName
	terraformDirPath := initPath + "/" + TerraformDir
	terraformStatePath := fmt.Sprintf("%s/%s/%s_%s_%s.tfstate", wp, prefix, wp, prefix, initDir)

	tCfg := fmt.Sprintf(
		terragruntConfigTempl,
		c.String(flags.S3StateBucket),
		terraformStatePath,
		c.String(flags.Region),
		c.String(flags.DynamoDBLockTable),
		c.String(flags.AuditProfile),
	)

	if err := os.RemoveAll(terraformDirPath); !os.IsNotExist(err) && err != nil {
		return fmt.Errorf("delete .terraform error: %+v", err)
	}

	if err := os.RemoveAll(terragruntConfigPath); !os.IsNotExist(err) && err != nil {
		return fmt.Errorf("delete %s error: %+v", TerragruntConfigName, err)
	}

	err := ioutil.WriteFile(terragruntConfigPath, []byte(tCfg), 0777)
	if err != nil {
		return fmt.Errorf("creating terragrunt config error: %+v", err)
	}

	return execute([]string{"init", "--terragrunt-config", terragruntConfigPath}, c, opts...)
}

func getPaths(path string) (string, string) {
	if path != "" {
		initPath := strings.Split(path, "/")
		return initPath[len(initPath)-1], path
	}

	return CurrentDir(), currentPath()
}
