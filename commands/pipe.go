package commands

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"trdeploy/flags"
)

type pipelineSteps struct {
	Steps []step `yaml:"steps"`
}

type step struct {
	Name    string `yaml:"name"`
	Threads []thread `yaml:"threads"`
}

type thread struct {
	Name       string `yaml:"name"`
	Path       string `yaml:"path"`
	VarProfile string `yaml:"var_profile"`
}


func parsePipeYaml(pipelineFile string) (*pipelineSteps, error) {
	if _, err := os.Stat(pipelineFile); err != nil {
		return nil, fmt.Errorf("file does not exist: '%s'. %s", pipelineFile, err)
	}
	stepsBytes, err := ioutil.ReadFile(pipelineFile)
	if err != nil {
		return nil, fmt.Errorf("cannot read from file: '%s'. %s", pipelineFile, err)
	}

	var steps pipelineSteps

	err = yaml.Unmarshal(stepsBytes, &steps)
	return &steps, err
}

func pipe(c *cli.Context, action func(thread, ...CommandOption) error) error {
	multithread := c.IsSet(flags.Multithread) && c.Bool(flags.Multithread)
	steps, ok := c.Value(stepsCtx).(*pipelineSteps)
	if !ok {
		return fmt.Errorf("not pipeline steps: %+v", c.Value(stepsCtx))
	}
	var wg sync.WaitGroup

	exit := make(chan os.Signal, 1)
	signal.Notify(exit, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	Steps:
		for i, s := range steps.Steps {
			errs := make(chan error, len(s.Threads))

		Threads:
			for j, th := range s.Threads {
				opt := ThreadName(s.Name, th.Name)
				wg.Add(1)

				log.Debugf("\n*** Step %d %s, Thread %d %s\n", i+1, s.Name, j+1, th.Name)

				if !multithread {
					if err := action(th, opt); err != nil {
						errs <- err
						wg.Done()
						break Threads
					}

					wg.Done()
					continue Threads
				}

				go func(th thread) {
					defer wg.Done()
					errs <- action(th, opt)
				}(th)
			}

			wg.Wait()

			if isExited(errs, exit) {
				break Steps
			}
		}

	return nil
}

func isExited(errs chan error, exit chan os.Signal) bool {
	go func(errs chan error) { close(errs) }(errs)

	var exited bool

	for err := range errs {
		if err != nil {
			log.Error(err)
			exited = true
		}
	}

	select {
	case <-exit:
		signal.Stop(exit)
		exited = true
	default:
		return exited
	}

	return exited
}
