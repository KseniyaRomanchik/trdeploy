package commands

import (
	"context"
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"sync"
	"trdeploy/flags"
)

type pipelineSteps struct {
	Steps []step `yaml:"steps"`
}

type step struct {
	Name    string
	Threads []thread
}

type thread struct {
	Name       string `yaml:"name"`
	Path       string `yaml:"path"`
	VarProfile string `yaml:"var_profile"`
}

func loadSteps(reverse bool) func (c *cli.Context) error {
	return func(c *cli.Context) error {
		pipelineFile := fmt.Sprintf("%s/%s", c.String(flags.GlobalPiplineProfile), c.String(flags.PiplineFile))

		steps, err := parsePipeYaml(pipelineFile)
		if err != nil {
			return err
		}

		if reverse {
			for i, j := 0, len(steps.Steps)-1; i < j; i, j = i+1, j-1 {
				steps.Steps[i], steps.Steps[j] = steps.Steps[j], steps.Steps[i]
			}
		}

		c.Context = context.WithValue(c.Context, stepsCtx, steps)

		return nil
	}
}


func parsePipeYaml(pipelineFile string) (*pipelineSteps, error) {
	if _, err := os.Stat(pipelineFile); err != nil {
		return nil, fmt.Errorf("File does not exist: '%s'. %s", pipelineFile, err)
	}
	stepsBytes, err := ioutil.ReadFile(pipelineFile)
	if err != nil {
		return nil, fmt.Errorf("Cannot read from file: '%s'. %s", pipelineFile, err)
	}

	var steps pipelineSteps

	err = yaml.Unmarshal(stepsBytes, &steps)
	return &steps, err
}

func pipe(c *cli.Context, action func(th thread) error) error {
		multithread := c.IsSet(flags.Multithread) && c.Bool(flags.Multithread)
		steps, ok := c.Value(stepsCtx).(*pipelineSteps)
		if !ok {
			return fmt.Errorf("not pipline steps: %+v", c.Context.Value("steps"))
		}
		var wg sync.WaitGroup

	Steps:
		for i, s := range steps.Steps {
			wg.Add(len(s.Threads))

			errs := make(chan error, len(s.Threads))
		Threads:
			for j, th := range s.Threads {
				log.Debugf("\n*** Step %d %s, Thread %d %s\n", i+1, s.Name, j+1, th.Name)
				th.Name = s.Name + " " + th.Name

				if !multithread {
					errs <- action(th)
					wg.Done()
					continue Threads
				}

				go func(th thread) {
					defer wg.Done()

					errs <- action(th)
				}(th)
			}

			wg.Wait()

			go func() {
				close(errs)
			}()

			for err := range errs {
				if err != nil {
					break Steps
				}
			}
		}

		return nil
	}