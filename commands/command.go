package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"os/exec"
	"strconv"
	"trdeploy/flags"
)

type CommandOption func(Command) Command
type Command struct {
	*exec.Cmd
	ThreadName []string
}

func Dir(dir string) CommandOption {
	return func(c Command) Command {
		c.Dir = dir
		return c
	}
}

func ThreadName(names ...string) CommandOption {
	return func(c Command) Command {
		if c.ThreadName == nil {
			c.ThreadName = names
			return c
		}

		c.ThreadName = append(c.ThreadName, names...)
		return c
	}
}

func AutoApprove() CommandOption {
	return func(c Command) Command {
		c.Args = append(c.Args, "-auto-approve")
		return c
	}
}

func Parallelism(v int) CommandOption {
	return func(c Command) Command {
		c.Args = append(c.Args, "-parallelism", strconv.Itoa(v))
		return c
	}
}

func Cmd(ctx *cli.Context) CommandOption {
	prefix := ctx.String(flags.Prefix)
	ap := ctx.String(flags.AuditProfile)
	wp := ctx.String(flags.WorkProfile)
	gvp := ctx.String(flags.GlobalVarPath)
	mtfv := ctx.String(flags.ModuleTfvars)

	return func(c Command) Command {
		c.Args = append(c.Args,
			"-var-file", fmt.Sprintf("%s/common.tfvars", gvp),
			"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
			"-var-file", mtfv,
			"-var", fmt.Sprintf("prefix=%s", prefix),
			"-var", fmt.Sprintf("aws_audit=%s", ap),)
		return c
	}
}
