package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

var (
	port       = flag.Int("port", 8765, "the port to listen on")
	remoteHost = flag.String("remote-host", "localhost", "the remote host to connect to")
	remotePort = flag.String("remote-port", 9876, "the remote host to connect to")
)

const buffSize int = 1024

func handleError(err error, msg string, args ...interface{}) {
	if msg == "" {
		msg = "Error"
	}

	if len(args) > 0 {
		msg = fmt.Sprintf(msg, args...)
	}

	fmt.Printf("%s: %v", msg, err)
	os.Exit(1)
}

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		handleError(err, "Error listening on port %v", *port)
	}

	con, err := ln.Accept()
	if err != nil {
		handleError(err, "Error accepting connection")
	}

	b := make([]byte, buffSize, buffSize)

	var count int

	for {
		count, err = con.Read(b)
		if err == io.EOF {
			os.Exit(0)
		}
		if err != nil {
			handleError(err, "Error reading from socket")
		}
		b2 := b[0:count]
		fmt.Print(string(b2))

		_, err = con.Write(b2)
		if err != nil {
			handleError(err, "Error writing to socket")
		}
	}
}
