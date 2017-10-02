package main

import (
	"encoding/json"
	"github.com/pkg/errors"
	"log"
	"net"
	"strings"
)

type Relay struct {
	control net.Conn
	listen  net.Listener
	conns   []*relayConn
}

type relayConn struct {
	clientCon  *net.Conn
	serverConn *net.Conn
}

func NewRelay(control net.Conn) *Relay {
	return &Relay{control: control}
}

func (r *Relay) Start() error {
	ln, err := net.Listen("tcp", ":0")
	if err != nil {
		return errors.Wrap(err, "Error creating listening socket")
	}
	r.listen = ln

	addr := r.listen.Addr().String()
	idx := strings.LastIndex(addr, ":")

	ip := ""
	port := addr

	if idx != -1 {
		ip = addr[:idx]
		port = addr[idx+1:]
	}

	data := map[string]interface{}{}
	data["ip"] = ip
	data["port"] = port

	bytes, err := json.Marshal(&data)
	if err != nil {
		return errors.Wrap(err, "Error marshaling data")
	}

	_, err = r.control.Write(bytes)
	if err != nil {
		return errors.Wrap(err, "Error sending data")
	}

	log.Printf("Creating new relay on %s", addr)

	for {
		con, err := r.listen.Accept()
		if err != nil {
			log.Printf("Error accepting socket: %s", err)
			continue
		}

		rc := &relayConn{clientCon: &con}
		r.conns = append(r.conns, rc)
	}
}

func (rc *relayConn) Start() {

}
