package command

import (
	"fmt"
	"flag"

	"github.com/mitchellh/colorstring"
)

// Agent is a Command implementation that starts the agent process
type Agent struct {
}

type AgentConfig struct {
	Server string
	NewrelicKey string
}

// Run is awesome
func (a *Agent) Run(args []string) int {
	var config AgentConfig

        cmdFlags := flag.NewFlagSet("agent", flag.ContinueOnError)
	configFile := cmdFlags.String("config", "./newrelic-mysql-agent.json", "Path to the agent configuration file")
	cmdFlags.StringVar(&config.Server, "server", "localhost", "IP or Hostname of MySQL server")
	cmdFlags.StringVar(&config.NewrelicKey, "newrelic-key", "", "The NewRelic license key to use for reporting")
	cmdFlags.Parse(args)
	fmt.Println(colorstring.Color("[green]B[red]O[cyan]O"), config.Server, config.NewrelicKey, *configFile)
	return 0
}

// Help is awesome
func (a *Agent) Help() string {
	return "starts the agent monitoring"
}

// Synopsis is awesome
func (a *Agent) Synopsis() string {
	return "connects to the MySQL server, collects statistics, and reports to NewRelic"
}

// ParseConfig parses the supplied configuration file and sets up the
// configuration structure
