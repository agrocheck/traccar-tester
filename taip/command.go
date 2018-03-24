package taip

import (
	"context"
	"flag"

	"github.com/google/subcommands"
)

type Command struct {
	id string
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
}

func (cmd *Command) Execute(_ context.Context, _ *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	// TODO: implement
	return subcommands.ExitSuccess
}
