package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
	"io"
	"log"
	"os"
	"os/exec"
	"trdeploy/flags"
)

var (
	Commands []*cli.Command
)

func LoadCommands() {
	Commands = []*cli.Command{
		{
			Name:      Init,
			Aliases:   []string{"i"},
			UsageText: "*** init ",
			Usage:     "init",
			Before:    beforeAction,
			Flags:     flags.Flags,
			Action:    commandAction(initAction),
		},
		{
			Name:      Plan,
			Aliases:   []string{"p"},
			UsageText: "*** plan ",
			Usage:     "plan",
			Before:    beforeAction,
			Flags:     flags.Flags,
			Action:    commandAction(initAction, plan),
		},
		{
			Name:      Apply,
			Aliases:   []string{"a"},
			UsageText: "*** apply ",
			Usage:     "apply",
			Before:    beforeAction,
			Flags:     flags.Flags,
			Action:    commandAction(initAction, apply),
		},
		{
			Name:      Destroy,
			Aliases:   []string{"d"},
			UsageText: "*** destroy ",
			Usage:     "destroy",
			Before:    beforeAction,
			Flags:     flags.Flags,
			Action:    commandAction(initAction, destroy),
		},
	}
}

func commandAction(actionFns ...func(c *cli.Context) error) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		fmt.Printf("\n %s %s\n", c.Command.UsageText, CurrentDir())
		for _, f := range c.App.VisibleFlags() {
			fmt.Printf("\t *  %s: %+v\n", f.Names()[0], c.String(f.Names()[0]))
		}

		for _, fn := range actionFns {
			if err := fn(c); err != nil {
				return err
			}
		}

		return nil
	}
}

func beforeAction(c *cli.Context) error {
	replaceModuleTfvars(c)

	return altsrc.InitInputSourceWithContext(flags.Flags, altsrc.NewYamlSourceFromFlagFunc("config"))(c)
}

func replaceModuleTfvars(c *cli.Context) {
	var newMtv string

	if !c.IsSet(flags.ModuleTfvars) {
		wp := c.String(flags.WorkProfile)
		newMtv = fmt.Sprintf("var/%s.tfvars", wp)
	} else {
		mtv := c.String(flags.ModuleTfvars)
		newMtv = fmt.Sprintf("var/%s", mtv)
	}

	c.Set(flags.ModuleTfvars, newMtv)
}

func command(args []string, c *cli.Context) error {
	if c.IsSet(flags.AdditionalArgs) {
		args = append(args, c.String(flags.AdditionalArgs))
	}

	cmd := exec.Command("terragrunt", args...)
	fmt.Printf("\n[command]: %s \n\n", cmd.String())

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		cli.Exit(fmt.Sprintf("terragrunt %s error: %s", args[0], err), 1)
		return fmt.Errorf("\nterragrunt %s error: %s", args[0], err)
	}

	if c.IsSet(flags.OutPlanLog) {
		logFile, err := os.Create(c.String(flags.OutPlanLog))
		if err != nil {
			cli.Exit(fmt.Sprintf("creating out plan log file error: %s", err), 1)
			return fmt.Errorf("creating out plan log file error: %s", err)
		}
		defer logFile.Close()

		wrt := io.MultiWriter(os.Stdout, logFile)
		log.SetOutput(wrt)
	}

	return nil
}
