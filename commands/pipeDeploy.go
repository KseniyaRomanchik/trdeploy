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
	ThreadName string `yaml:"thread_name"`
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

	for i, step := range steps {
		wg.Add(len(step))

		for j, thread := range step {
			log.Debugf("\n*** Step %d, Thread %d\n", i, j)

			if !multithread {
				deploy(thread, c, multithread)
				wg.Done()
				continue
			}

			go func(thread pipelineSteps) {
				log.Debugf("*** Goroutine %s\n", thread.ThreadName)
				defer wg.Done()

				deploy(thread, c, multithread)
			}(thread)
		}

		wg.Wait()
	}

	return nil
}

func parsePipeYaml(pipelineFile string) ([][]pipelineSteps, error) {
	if _, err := os.Stat(pipelineFile); err != nil {
		return nil, fmt.Errorf("File does not exist: '%s'. %s", pipelineFile, err)
	}
	stepsBytes, err := ioutil.ReadFile(pipelineFile)
	if err != nil {
		return nil, fmt.Errorf("Cannot read from file: '%s'. %s", pipelineFile, err)
	}

	var steps [][]pipelineSteps

	err = yaml.Unmarshal(stepsBytes, &steps)
	return steps, err
}

func deploy(thread pipelineSteps, c *cli.Context, multithread bool) {
	opts := []CommandOption{Dir(thread.Path), Env([]string{thread.ThreadName})}

	if err := initAction(thread.Path)(c, opts...); err != nil {
		log.Errorf("%s, Init pipe-deploy error %s: %s", thread.ThreadName, thread.Path, err)
		return
	}

	if multithread {
		opts = append(opts, AutoApprove())
	}

	if err := apply(c, opts...); err != nil {
		log.Errorf("%s, Apply pipe-deploy error %s: %s", thread.ThreadName, thread.Path, err)
		return
	}
}
