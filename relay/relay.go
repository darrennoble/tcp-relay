package main

import (
	"flag"
	"fmt"
	"github.com/darrennoble/tcp-utils/util"
	"io"
	"net"
	"os"
)

var (
	port       = flag.Int("port", 8765, "the port to listen on")
	remoteHost = flag.String("remote-host", "localhost", "the remote host to connect to")
	remotePort = flag.Int("remote-port", 9876, "the remote host to connect to")
)

const buffSize int = 1024

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		util.HandleError(err, "Error listening on port %v", *port)
	}

	con, err := ln.Accept()
	if err != nil {
		util.HandleError(err, "Error accepting connection")
	}

	b := make([]byte, buffSize, buffSize)

	var count int

	for {
		count, err = con.Read(b)
		if err == io.EOF {
			os.Exit(0)
		}
		if err != nil {
			util.HandleError(err, "Error reading from socket")
		}
		b2 := b[0:count]
		fmt.Print(string(b2))

		_, err = con.Write(b2)
		if err != nil {
			util.HandleError(err, "Error writing to socket")
		}
	}
}
