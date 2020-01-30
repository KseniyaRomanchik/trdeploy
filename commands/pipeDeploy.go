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
	gpp := c.String(flags.GlobalPiplineProfile)

	steps, err := parsePipeYaml(gpp)
	if err != nil {
		return err
	}

	var wg sync.WaitGroup

	for i, step := range steps {
		wg.Add(len(step))

		for j, thread := range step {
			log.Debugf("\n*** Step %d, Thread %d\n", i, j)

			go func(thread pipelineSteps) {
				log.Debugf("*** Goroutine %s\n", thread.ThreadName)
				defer wg.Done()

				deploy(thread, c)
			}(thread)
		}

		wg.Wait()
	}

	return nil
}

func parsePipeYaml(gpp string) ([][]pipelineSteps, error) {
	if _, err := os.Stat(gpp); err != nil {
		return nil, fmt.Errorf("File is not exist: '%s'. %s", gpp, err)
	}
	stepsBytes, err := ioutil.ReadFile(gpp)
	if err != nil {
		return nil, fmt.Errorf("Cannot read from file: '%s'. %s", gpp, err)
	}

	var steps [][]pipelineSteps

	err = yaml.Unmarshal(stepsBytes, &steps)
	return steps, err
}

func deploy(thread pipelineSteps, c *cli.Context) {
	if err := initAction(c, Dir(thread.Path)); err != nil {
		fmt.Printf("*** Goroutine %s, Init pipe-deploy error %s: %s", thread.ThreadName, thread.Path, err)
		return
	}

	if err := apply(c, Dir(thread.Path)); err != nil {
		fmt.Printf("*** Goroutine %s, Apply pipe-deploy error %s: %s", thread.ThreadName, thread.Path, err)
		return
	}
}
