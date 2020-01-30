package commands

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"os/exec"
	"trdeploy/flags"
)

type CommandOption func(*exec.Cmd) *exec.Cmd

func Dir(dir string) CommandOption {
	return func(c *exec.Cmd) *exec.Cmd {
		c.Dir = dir
		return c
	}
}

func execute(args []string, c *cli.Context, opts ...CommandOption) error {
	if c.IsSet(flags.AdditionalArgs) {
		args = append(args, c.String(flags.AdditionalArgs))
	}

	args = append(args, "--terragrunt-config", TerragruntConfigPath())

	cmd := exec.Command("terragrunt", args...)

	for _, opt := range opts {
		cmd = opt(cmd)
	}

	log.Infof("[command]: %s in %s\n\n", cmd.String(), cmd.Dir)

	cmd.Stdout = os.Stdout
	cmd.Stdin = os.Stdin
	cmd.Stderr = os.Stderr

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
