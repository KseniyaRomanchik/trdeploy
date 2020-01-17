package main

import (
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"trdeploy/commands"
	"trdeploy/flags"
)

func main() {
	err := flags.LoadFlags()
	if err != nil {
		log.Fatal(err)
	}

	commands.LoadCommands()

	app := &cli.App{
		Name:     "trdeploy",
		Usage:    "./trdeploy {command} --work_profile {aws} --prefix {prefix}  --arg1 value1  --arg2 value2",
		Version:  "0.0.1",
		Flags:    flags.Flags,
		Commands: commands.Commands,
	}

	err = app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
