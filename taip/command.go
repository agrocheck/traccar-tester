package taip

import (
	"context"
	"flag"
	"log"

	"github.com/agrocheck/traccar-tester"
	"github.com/google/subcommands"
)

type Command struct {
	id string
	event string
	count int
	backoff int
}

func (cmd *Command) Name() string {
	return "taip"
}

func (cmd *Command) Synopsis() string {
	return "TODO"
}

func (cmd *Command) Usage() string {
	return "TODO"
}

func (cmd *Command) SetFlags(f *flag.FlagSet) {
	f.StringVar(&cmd.id, "id", "", "TODO")
	f.StringVar(&cmd.event, "event", "", "TODO")
	f.IntVar(&cmd.count, "count", 1, "TODO")
	f.IntVar(&cmd.backoff, "backoff", 0, "TODO")
}

func (cmd *Command) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// TODO: input validations

	if len(cmd.id) == 0 {
		log.Print("id not defined")
		return subcommands.ExitFailure
	}

	if len(cmd.id) != 4 {
		log.Print("id must be a 4-length string")
		return subcommands.ExitFailure
	}

	p := NewDevice(cmd.id, cmd.event)

	return tester.Run(p, cmd.count, cmd.backoff)
}
