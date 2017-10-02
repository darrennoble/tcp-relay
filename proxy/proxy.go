package main

import (
	"flag"
	"fmt"
	"github.com/darrennoble/tcp-utils/errors"
	"io"
	"net"
)

var (
	port       = flag.Int("port", 8765, "the port to listen on")
	remoteHost = flag.String("remote-host", "localhost", "the remote host to connect to")
	remotePort = flag.Int("remote-port", 9876, "the remote host to connect to")
)

func main() {
	flag.Parse()

	ln, err := net.Listen("tcp", fmt.Sprintf(":%v", *port))
	if err != nil {
		errors.Fatal(err, "Error listening on port %v", *port)
	}
	defer ln.Close()

	proxyHostStr := fmt.Sprintf("%s:%v", *remoteHost, *remotePort)

	for {
		clientCon, err := ln.Accept()
		if err != nil {
			errors.Print(err, "Error accepting connection")
			continue
		}

		proxyCon, err := net.Dial("tcp", proxyHostStr)
		if err != nil {
			errors.Print(err, "Error connecting to proxy host %s", proxyHostStr)
		}

		go copyData(clientCon, proxyCon)
		go copyData(proxyCon, clientCon)
	}
}

func copyData(src, dest net.Conn) {
	io.Copy(src, dest)
	src.Close()
	dest.Close()
}
