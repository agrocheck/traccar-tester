package tester

import (
	"github.com/agrocheck/traccar-tester/taip"
	"github.com/google/subcommands"
)

func Commands() (cmds []subcommands.Command) {
	cmds = append(cmds, &taip.Command{})
	return
}
