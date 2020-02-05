package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"os"
	"strings"
	"trdeploy/flags"
)

func initAction(args ...string) func(*cli.Context, ...CommandOption) error {
	return func(c *cli.Context, opts ...CommandOption) error {
		currentDir := CurrentDir()
		path := "./"

		if len(args) > 0 {
			path = args[0] + "/"
			currentPath := strings.Split(args[0], "/")
			currentDir = currentPath[len(currentPath)-1]
		}

		if err := os.RemoveAll(path + ".terraform"); !os.IsNotExist(err) && err != nil {
			cli.Exit(fmt.Sprintf("delete .terraform error: %+v", err), 1)
			return fmt.Errorf("delete .terraform error: %+v", err)
		}

		if err := os.RemoveAll(path + terragruntConfigName); !os.IsNotExist(err) && err != nil {
			cli.Exit(fmt.Sprintf("delete %s error: %+v", terragruntConfigName, err), 1)
			return fmt.Errorf("delete %s error: %+v", terragruntConfigName, err)
		}

		prefix := c.String(flags.Prefix)
		wp := c.String(flags.WorkProfile)

		tCfg := fmt.Sprintf(
			terragruntConfigTempl,
			c.String(flags.S3StateBacket),
			fmt.Sprintf("%s/%s/%s_%s_%s.tfstate", wp, prefix, wp, prefix, currentDir),
			c.String(flags.Region),
			c.String(flags.DynamodbLockTable),
			c.String(flags.AuditProfile),
		)

		err := ioutil.WriteFile(path+"/"+terragruntConfigName, []byte(tCfg), 0777)
		if err != nil {
			cli.Exit(fmt.Sprintf("creating terragrunt config error: %+v", err), 1)
			return fmt.Errorf("creating terragrunt config error: %+v", err)
		}

		return execute([]string{"init", "--terragrunt-config", path+"/"+terragruntConfigName}, c, opts...)
	}
}
