package main

import (
	"flag"
	"fmt"
	"github.com/darrennoble/tcp-utils/errors"
	"io"
	"net"
)

var (
	port = flag.Int("port", 9876, "the port to listen on")
)

const buffSize int = 1024

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
		}

		go func(c net.Conn) {
			io.Copy(c, c)
			c.Close()
		}(con)
	}
}
