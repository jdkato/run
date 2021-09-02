package cli

import (
	"flag"
)

var commandInfo = map[string]string{
	"ls-config": "Print the current configuration to stdout and exit.",
}

// Actions are the available CLI commands.
var Actions = map[string]func(args []string) error{
	"help": printUsage,
}

func printUsage(args []string) error {
	flag.Usage()
	return nil
}
