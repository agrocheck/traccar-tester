package taip

import (
	"context"
	"flag"
	"log"

	"github.com/agrocheck/traccar-tester"
	"github.com/google/subcommands"
)

type Command struct {
	traccar tester.Traccar
	id      string
	event   string
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
	// TODO: move these to tester.go?
	f.StringVar(&cmd.traccar.Host, "host", "localhost", "TODO")
	f.StringVar(&cmd.traccar.Port, "port", "5031", "TODO")
	f.StringVar(&cmd.traccar.Network, "network", "tcp", "TODO")
	f.IntVar(&cmd.traccar.Count, "count", 1, "TODO")
	f.IntVar(&cmd.traccar.Backoff, "backoff", 0, "TODO")

	f.StringVar(&cmd.id, "id", "", "TODO")
	f.StringVar(&cmd.event, "event", "", "TODO")
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

	if len(cmd.event) == 0 {
		log.Print("event not defined")
		return subcommands.ExitFailure
	}

	if len(cmd.event) != 2 {
		log.Print("event must be a 2-length string")
		return subcommands.ExitFailure
	}

	p := NewDevice(cmd.id, cmd.event)

	return tester.Run(&cmd.traccar, p)
}
