package flags

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"os"
)

var (
	Flags []cli.Flag
)

func LoadFlags() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := fmt.Sprintf("%s/.%s.yaml", home, "trdeploy")
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configPath = fmt.Sprintf("/etc/%s.yaml", "trdeploy")
	}

	Flags = []cli.Flag{
		&cli.StringFlag{
			Name:  Region,
			Usage: "aws region",
			Value: "us-west-2",
		},
		&cli.StringFlag{
			Name:  DynamodbLockTable,
			Usage: "terraform state-lock",
			Value: "unitedsoft-terraform-state-backet-lock",
		},
		&cli.StringFlag{
			Name:  S3StateBacket,
			Usage: "terraform s3 state backet",
			Value: "unitedsoft-terraform-state-backet",
		},
		&cli.StringFlag{
			Name:  AuditProfile,
			Usage: "aws-audit-profile",
			Value: "default",
		},
		&cli.StringFlag{
			Name:  OutPlanLog,
			Usage: "out-plan-log",
		},
		&cli.StringFlag{
			Name:  AdditionalArgs,
			Usage: "additional-args",
		},
		&cli.StringFlag{
			Name:  ModuleTfvars,
			Usage: "name of module ftvars-file  (default  {aws-profile}.tfvars)",
		},
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  GlobalVarPath,
			Usage: "path  global var.tf (default from /etc/tdeploy.yaml)",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  GlobalPiplineProfile,
			Usage: "path  global var.tf (default from /etc/tdeploy.yaml)",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  BasePath,
			Usage: "path  base var.tf (default from /etc/tdeploy.yaml)",
		}),
		&cli.StringFlag{
			Name:  Config,
			Value: configPath,
		},
		&cli.StringFlag{
			Name:  PlanFile,
			Usage: "plan-file",
		},
		&cli.StringFlag{
			Name:  DeployProfile,
			Usage: "deploy-profile",
		},
		&cli.StringFlag{
			Name:  Prefix,
			Usage: "prefix",
		},
		&cli.StringFlag{
			Name:  WorkProfile,
			Usage: "work-profile",
		},
	}

	return nil
}
