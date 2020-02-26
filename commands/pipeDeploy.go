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

	for i, s := range steps.Steps {
		wg.Add(len(s.Threads))

		for j, th := range s.Threads {
			log.Debugf("\n*** Step %d %s, Thread %d %s\n", i+1, s.Name, j+1, th.Name)
			th.Name = s.Name + " " + th.Name

			if !multithread {
				deploy(th, c, multithread)
				wg.Done()
				continue
			}

			go func(th thread) {
				defer wg.Done()

				deploy(th, c, multithread)
			}(th)
		}

		wg.Wait()
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

func deploy(th thread, c *cli.Context, multithread bool) {
	opts := []CommandOption{Dir(th.Path), Env([]string{th.Name})}
	c.Set(flags.ExecDir, th.Path)

	if err := initAction(c, opts...); err != nil {
		log.Errorf("%s, Init pipe-deploy error %s: %s", th.Name, th.Path, err)
		return
	}

	if multithread {
		opts = append(opts, AutoApprove())
	}

	if err := apply(c, opts...); err != nil {
		log.Errorf("%s, Apply pipe-deploy error %s: %s", th.Name, th.Path, err)
		return
	}
}
