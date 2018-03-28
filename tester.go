package tester

import (
	"fmt"
	"log"
	"net"
	"time"

	"github.com/google/subcommands"
)

type Traccar struct {
	Host    string
	Port    string
	Network string
	Count   int
	Backoff int
}

type Protocol interface {
	RandomMessage() string
}

func Run(traccar *Traccar, p Protocol) subcommands.ExitStatus {
	address := traccar.Host + ":" + traccar.Port
	conn, err := net.Dial(traccar.Network, address)

	if err != nil {
		log.Print(err)
		return subcommands.ExitFailure
	}

	for i := 0; i < traccar.Count; i++ {
		fmt.Fprintln(conn, p.RandomMessage())
		time.Sleep(time.Duration(traccar.Backoff) * time.Second)
	}

	return subcommands.ExitSuccess
}
