package commands

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"os/exec"
	"strconv"
	"trdeploy/flags"
)

type CommandOption func(*exec.Cmd) *exec.Cmd

func Dir(dir string) CommandOption {
	return func(c *exec.Cmd) *exec.Cmd {
		c.Dir = dir
		return c
	}
}

func Env(env []string) CommandOption {
	return func(c *exec.Cmd) *exec.Cmd {
		if c.Env == nil {
			c.Env = append(env, os.Environ()...)
		}

		c.Env = append(env, c.Env...)
		return c
	}
}

func AutoApprove() CommandOption {
	return func(c *exec.Cmd) *exec.Cmd {
		c.Args = append(c.Args, "-auto-approve")
		return c
	}
}

func Parallelism(v int) CommandOption {
	return func(c *exec.Cmd) *exec.Cmd {
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

	return func(c *exec.Cmd) *exec.Cmd {
		c.Args = append(c.Args,
			"-var-file", fmt.Sprintf("%s/common.tfvars", gvp),
			"-var-file", fmt.Sprintf("%s/%s.tfvars", gvp, wp),
			"-var-file", mtfv,
			"-var", fmt.Sprintf("prefix=%s", prefix),
			"-var", fmt.Sprintf("aws_audit=%s", ap),)
		return c
	}
}

func execute(args []string, c *cli.Context, opts ...CommandOption) error {
	if c.IsSet(flags.AdditionalArgs) {
		args = append(args, c.String(flags.AdditionalArgs))
	}

	cmd := exec.Command("terragrunt", args...)

	for _, opt := range opts {
		cmd = opt(cmd)
	}

	log.Debugf("[command]: %s in %s\n\n", cmd.String(), cmd.Dir)

	cmd.Stdin = os.Stdin

	if cmd.Env == nil || len(cmd.Env) == 0 {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		printThreadOutput(cmd)
	}

	if err := cmd.Run(); err != nil {
		cli.Exit(fmt.Sprintf("terragrunt %s error: %s", args[0], err), 1)
		return fmt.Errorf("terragrunt %s error: %s", args[0], err)
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

func printThreadOutput(cmd *exec.Cmd) {
	threadName := cmd.Env[0]

	outReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Errorln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		cli.Exit(err, 1)
	}

	outScanner := bufio.NewScanner(outReader)
	go func() {
		defer outReader.Close()
		for outScanner.Scan() {
			log.Infof("%s | %s\n", threadName, outScanner.Text())
		}
	}()

	errReader, err := cmd.StderrPipe()
	if err != nil {
		log.Errorln(os.Stderr, "Error creating StderrPipe for Cmd", err)
		cli.Exit(err, 1)
	}

	errScanner := bufio.NewScanner(errReader)
	go func() {
		defer errReader.Close()
		for errScanner.Scan() {
			log.Infof("%s | %s\n", threadName, errScanner.Text())
		}
	}()
}
