package tester

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"time"

	"github.com/google/subcommands"
)

type Protocol interface {
	RandomMessage() string
}

func Run(p Protocol, count int, backoff int) subcommands.ExitStatus {
	// TODO: variables
	host := "localhost"
	port := 8080
	network := "tcp"

	address := host + ":" + strconv.Itoa(port)
	conn, err := net.Dial(network, address)

	if err != nil {
		log.Print(err)
		return subcommands.ExitFailure
	}

	for i := 0; i < count; i++ {
		fmt.Fprintln(conn, p.RandomMessage())
		time.Sleep(time.Duration(backoff) * time.Second)
	}

	return subcommands.ExitSuccess
}
