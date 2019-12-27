package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"trdeploy/flags"
)

var Commands = []*cli.Command{
	{
		Name:      "init",
		Aliases:   []string{"i"},
		UsageText: "*** init ",
		Usage:     "init",
		Flags:     flags.Flags,
		Action:    commandAction(InitAction),
	},
	{
		Name:      "plan",
		Aliases:   []string{"p"},
		UsageText: "*** plan ",
		Usage:     "plan",
		Flags:     flags.Flags,
		Action:    commandAction(InitAction, Plan),
	},
	{
		Name:      "apply",
		Aliases:   []string{"a"},
		UsageText: "*** apply ",
		Usage:     "apply",
		Flags:     flags.Flags,
		Action:    commandAction(Apply),
	},
	{
		Name:      "destroy",
		Aliases:   []string{"d"},
		UsageText: "*** destroy ",
		Usage:     "destroy",
		Flags:     flags.Flags,
		Action:    commandAction(Destroy),
	},
}

func commandAction(actionFns ...func(c *cli.Context) error) func(c *cli.Context) error {
	return func(c *cli.Context) error {
		fmt.Printf("\n %s %s\n", c.Command.UsageText, CurrentDir())

		for _, fn := range actionFns {
			if err := fn(c); err != nil {
				return err
			}
		}

		return nil
	}
}
