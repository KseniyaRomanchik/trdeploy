package flags

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func printVersion(c *cli.Context) {
	fmt.Printf(versionTemplate, c.App.Version, Commit, Image, Time)
}
