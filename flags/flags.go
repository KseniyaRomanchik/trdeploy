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
	cli.VersionPrinter = printVersion

	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	configPath := fmt.Sprintf("%s/.%s.yaml", home, ConfigFileName)
	if _, err := os.Stat(configPath); os.IsNotExist(err) {
		configPath = fmt.Sprintf("/etc/%s.yaml", ConfigFileName)
	}

	Flags = []cli.Flag{
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  Region,
			Usage: "aws region",
			Value: "us-west-2",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  DynamoDBLockTable,
			Usage: "terraform state-lock",
			Value: "unitedsoft-terraform-state-backet-lock",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  S3StateBucket,
			Usage: "terraform s3 state bucket",
			Value: "unitedsoft-terraform-state-backet",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  AuditProfile,
			Usage: "aws-audit-profile",
			Value: "default",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  OutPlanLog,
			Usage: "out-plan-log",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  AdditionalArgs,
			Usage: "additional-args",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  ModuleTfvars,
			Usage: "name of module ftvars-file  (default  {aws-profile}.tfvars)",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  GlobalVarPath,
			Usage: "path  global var.tf (default from /etc/tdeploy.yaml)",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  GlobalPipelineProfile,
			Usage: "path  global var.tf (default from /etc/tdeploy.yaml)",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  BasePath,
			Usage: "path  base var.tf (default from /etc/tdeploy.yaml)",
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  Config,
			Value: configPath,
		}),
		altsrc.NewStringFlag(&cli.StringFlag{
			Name:  DeployProfile,
			Usage: "deploy-profile",
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:  Parallelism,
			Usage: "parallelism",
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:  Timeout,
			Usage: "timeout in seconds",
			Value: 10 * 60,
		}),
		altsrc.NewBoolFlag(&cli.BoolFlag{
			Name:  Test,
			Usage: "flag for functional tests",
		}),
		altsrc.NewIntFlag(&cli.IntFlag{
			Name:  LogLevel,
			Usage: "set log level",
			Value: 4,
		}),
	}

	return nil
}

var RequiredFlags = []cli.Flag{
	&cli.StringFlag{
		Name:     Prefix,
		Usage:    "prefix",
		Required: true,
	},
	&cli.StringFlag{
		Name:     WorkProfile,
		Usage:    "work-profile",
		Required: true,
	},
}

var PipeFlags = []cli.Flag{
	altsrc.NewBoolFlag(&cli.BoolFlag{
		Name:  Multithread,
		Usage: "multithread",
	}),
	altsrc.NewStringFlag(&cli.StringFlag{
		Name:     PipelineFile,
		Usage:    "Global pipeline profile file name",
		Required: true,
	}),
}

var ApplyFlags = []cli.Flag{
	altsrc.NewStringFlag(&cli.StringFlag{
		Name:  PlanFile,
		Usage: "plan-file",
	}),
}
