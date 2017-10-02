package main

import (
	"flag"
	"fmt"
	"github.com/darrennoble/tcp-utils/errors"
	"net"
)

var (
	port = flag.Int("port", 9999, "the port to request relays")
)

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		errors.Fatal(err, "Error listening on port %v", *port)
	}
	defer ln.Close()

	for {
		con, err := ln.Accept()
		if err != nil {
			errors.Print(err, "Error accepting connection")
			continue
		}

		relay := NewRelay(con)
		go relay.Start()
	}
}
