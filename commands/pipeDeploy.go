package commands

import (
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

func pipeDeploy(c *cli.Context, opts ...CommandOption) error {
	pipelineFile := fmt.Sprintf("%s/%s", c.String(flags.GlobalPiplineProfile), c.String(flags.PiplineFile))
	multithread := c.IsSet(flags.Multithread) && c.Bool(flags.Multithread)

	steps, err := parsePipeYaml(pipelineFile)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

Steps:
	for i, s := range steps.Steps {
		wg.Add(len(s.Threads))

		deployError := make(chan error, len(s.Threads))
	Threads:
		for j, th := range s.Threads {
			log.Debugf("\n*** Step %d %s, Thread %d %s\n", i+1, s.Name, j+1, th.Name)
			th.Name = s.Name + " " + th.Name

			if !multithread {
				deployError <- deploy(th, c, multithread)
				wg.Done()
				continue Threads
			}

			go func(th thread) {
				defer wg.Done()

				deployError <- deploy(th, c, multithread)
			}(th)
		}

		wg.Wait()

		go func() {
			close(deployError)
		}()

		for err := range deployError {
			if err != nil {
				break Steps
			}
		}
	}

	return nil
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

func deploy(th thread, c *cli.Context, multithread bool) error {
	bp := c.String(flags.BasePath)
	execPath := bp + "/" + th.Path
	opts := []CommandOption{Dir(execPath), Env([]string{th.Name})}

	c.Set(flags.ExecPath, execPath)

	if err := initAction(c, opts...); err != nil {
		log.Errorf("%s, Init pipe-deploy error %s: %s", th.Name, execPath, err)
		return err
	}

	if multithread {
		opts = append(opts, AutoApprove())
	}

	if err := apply(c, opts...); err != nil {
		log.Errorf("%s, Apply pipe-deploy error %s: %s", th.Name, execPath, err)
		return err
	}

	return nil
}
