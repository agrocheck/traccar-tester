package main

import (
	"context"
	"flag"
	"os"

	"github.com/agrocheck/traccar-tester"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	for _, cmd := range tester.Commands() {
		// TODO: check group
		subcommands.Register(cmd, "protocols")
	}

	// TODO: check
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
