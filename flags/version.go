package flags

import (
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
)

func printVersion(c *cli.Context) {
	log.Printf(versionTemplate, c.App.Version, Commit, Image, Time)
}
