package main

import (
	"log"
	"os"
	"github.com/econnell/newrelic-mysql-agent/command"
	"github.com/mitchellh/cli"
)

// Commands is the mapping of all the available Terraform commands.
var Commands map[string]cli.CommandFactory

// ErrorPrefix
const ErrorPrefix = "e:"

// OutputPrefix
const OutputPrefix = "o:"

func main() {
	c := cli.NewCLI("newrelic-mysql-agent", "1.0.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory{
		"agent": func() (cli.Command, error) {
			return &command.Agent{}, nil
		},
	}
	exitStatus, err := c.Run()
	if err != nil {
		log.Println(err)
	}
	os.Exit(exitStatus)
}
