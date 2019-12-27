package flags

import "github.com/urfave/cli/v2"

var Flags = []cli.Flag{
	&cli.StringFlag{
		Name:  "region",
		Usage: "aws region",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "dynamodb_lock_table",
		Usage: "terraform state_lock",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "s3_state_backet",
		Usage: "terraform s3 state backet",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "audit_profile",
		Usage: "aws_audit_profile",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "out_plan_log",
		Usage: "out_plan_log",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "additional_args",
		Usage: "additional_args",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "module_tfvars",
		Usage: "name of module ftvars-file  (default  {aws_profile}.tfvars)",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "global_var_path",
		Usage: "path  global var.tf (default from /etc/tdeploy.cnf)",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "plan_file",
		Usage: "plan_file",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "deploy_profile",
		Usage: "deploy_profile",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "prefix",
		Usage: "prefix",
		//Required: true,
	},
	&cli.StringFlag{
		Name:  "work_profile",
		Usage: "work_profile",
		//Required: true,
	},
}
