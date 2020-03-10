package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"trdeploy/flags"
)

func pipeDestroy(c *cli.Context, options ...CommandOption) error {
	return pipe(c, func(th thread, opts ...CommandOption) error {
		opts = append(opts, options...)

		bp := c.String(flags.BasePath)
		multithread := c.IsSet(flags.Multithread) && c.Bool(flags.Multithread)
		execPath := bp + "/" + th.Path
		opts = append(opts, Dir(execPath))

		if err := initAction(c, opts...); err != nil {
			return fmt.Errorf("%s, Init pipe-destroy error %s: %s", th.Name, execPath, err)
		}

		if multithread {
			opts = append(opts, AutoApprove())
		}

		if err := destroy(c, opts...); err != nil {
			return fmt.Errorf("%s, Apply pipe-destroy error %s: %s", th.Name, execPath, err)
		}

		return nil
	})
}
