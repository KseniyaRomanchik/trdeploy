package commands

import (
	"bufio"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"io"
	"os"
	"os/exec"
	"os/signal"
	"strings"
	"syscall"
	"trdeploy/flags"
)

func execute(args []string, c *cli.Context, opts ...CommandOption) error {
	if c.IsSet(flags.AdditionalArgs) {
		args = append(args, c.String(flags.AdditionalArgs))
	}

	cmd := Command{Cmd: exec.Command("terragrunt", args...)}

	for _, opt := range opts {
		cmd = opt(cmd)
	}

	log.Debugf("[command]: %s in %s\n\n", cmd.String(), cmd.Dir)

	cmd.Stdin = os.Stdin

	if cmd.ThreadName == nil || len(cmd.ThreadName) == 0 {
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr
	} else {
		printThreadOutput(cmd)
	}

	stopSignaling := signalingProcess(cmd)

	if err := cmd.Run(); err != nil {
		return fmt.Errorf("terragrunt %s error: %s", args[0], err)
	}

	stopSignaling()

	if c.IsSet(flags.OutPlanLog) {
		return savePlanLog(c.String(flags.OutPlanLog))
	}

	return nil
}

func printThreadOutput(cmd Command) {
	threadName := strings.Join(cmd.ThreadName, " ")

	outReader, err := cmd.StdoutPipe()
	if err != nil {
		log.Errorln(os.Stderr, "Error creating StdoutPipe for Cmd", err)
		return
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
		return
	}

	errScanner := bufio.NewScanner(errReader)
	go func() {
		defer errReader.Close()
		for errScanner.Scan() {
			log.Infof("%s | %s\n", threadName, errScanner.Text())
		}
	}()
}

func signalingProcess(cmd Command) func() {
	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func (cmd Command) {
		for sig := range exit {
			signal.Stop(exit)

			var threadMessage string
			if len(cmd.Env) > 0 {
				threadMessage = fmt.Sprintf("%s | ", cmd.Env[0])
			}

			log.Warnf("%sSignal interrupting message: %v", threadMessage, sig)

			if err := cmd.Process.Signal(sig); err != nil {
				log.Errorf("%sSignal interrupting error: %v", threadMessage, err)
			}
		}
	}(cmd)

	return func() { signal.Stop(exit) }
}

func savePlanLog(name string) error {
	logFile, err := os.Create(name)
	if err != nil {
		return fmt.Errorf("creating out plan log file error: %s", err)
	}
	defer logFile.Close()

	wrt := io.MultiWriter(os.Stdout, logFile)
	log.SetOutput(wrt)

	return nil
}
