package main

import (
	"context"
	"flag"
	"os"

	"github.com/agrocheck/traccar-tester/taip"
	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")
	// TODO: check group
	subcommands.Register(&taip.Command{}, "protocols")

	// TODO: check
	flag.Parse()
	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx)))
}
