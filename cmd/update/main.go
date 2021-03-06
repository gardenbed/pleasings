package main

import (
	"os"

	"github.com/gardenbed/charm/ui"
	"github.com/mitchellh/cli"

	"github.com/gardenbed/pleasings/internal/command/update"
	"github.com/gardenbed/pleasings/metadata"
)

func createCLI(ui ui.UI) *cli.CLI {
	c := cli.NewCLI("update", metadata.String())
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"": update.NewFactory(ui),
	}

	return c
}

func main() {
	u := ui.New(ui.Info)

	// Create the CLI app
	app := createCLI(u)

	code, err := app.Run()
	if err != nil {
		u.Errorf(ui.Red, "%s", err)
	}

	os.Exit(code)
}
