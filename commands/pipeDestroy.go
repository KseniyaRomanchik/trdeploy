package commands

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	"trdeploy/flags"
)

func pipeDestroy(c *cli.Context, options ...CommandOption) error {
	return pipe(c, func(th thread) error {
		opts := make([]CommandOption, len(options))
		copy(opts, options)

		bp := c.String(flags.BasePath)
		multithread := c.IsSet(flags.Multithread) && c.Bool(flags.Multithread)
		execPath := bp + "/" + th.Path
		opts = append(opts, []CommandOption{Dir(execPath), Env([]string{th.Name})}...)

		if err := initAction(c, opts...); err != nil {
			log.Errorf("%s, Init pipe-destroy error %s: %s", th.Name, execPath, err)
			return err
		}

		if multithread {
			opts = append(opts, AutoApprove())
		}

		if err := destroy(c, opts...); err != nil {
			log.Errorf("%s, Apply pipe-destroy error %s: %s", th.Name, execPath, err)
			return err
		}

		return nil
	})
}
