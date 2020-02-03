package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"github.com/urfave/cli/v2/altsrc"
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
		{
			Name:      PipeDeploy,
			UsageText: "*** pipe deploy ",
			Usage:     "pipe deploy",
			Before:    beforeAction,
			Flags:     flags.Flags,
			Action:    commandAction(pipeDeploy),
		},
	}
}

func commandAction(actionFns ...func(*cli.Context, ...CommandOption) error) func(c *cli.Context) error {
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

	mic, _ := altsrc.NewYamlSourceFromFlagFunc(configFileName)(c)

	return altsrc.InitInputSourceWithContext(
		c.App.Flags,
		func(ctx *cli.Context) (altsrc.InputSourceContext, error) {
			return prepareNestedInputSource(mic, c.String(flags.WorkProfile), c.App.Flags), nil
		})(c)
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
