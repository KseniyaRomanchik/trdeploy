package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"trdeploy/flags"
)

func pipeDeploy(c *cli.Context, options ...CommandOption) error {
	return pipe(c, func(th thread, opts ...CommandOption) error {
		opts = append(opts, options...)

		bp := c.String(flags.BasePath)
		multithread := c.IsSet(flags.Multithread) && c.Bool(flags.Multithread)
		execPath := bp + "/" + th.Path
		opts = append(opts, Dir(execPath))

		if err := initAction(c, opts...); err != nil {
			return fmt.Errorf("%s, Init pipe-deploy error %s: %s", th.Name, execPath, err)
		}

		if multithread {
			opts = append(opts, AutoApprove())
		}

		if !c.IsSet(flags.Parallelism) && th.Parallelism != 0 {
			opts = append(opts, Parallelism(th.Parallelism))
		}

		if err := apply(c, opts...); err != nil {
			return fmt.Errorf("%s, Apply pipe-deploy error %s: %s", th.Name, execPath, err)
		}

		return nil
	})
}
