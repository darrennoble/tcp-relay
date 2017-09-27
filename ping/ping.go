package main

import (
	"flag"
	"fmt"
	"github.com/darrennoble/tcp-utils/errors"
	"net"
	"time"
)

var (
	host        = flag.String("host", "localhost", "the host to send to")
	port        = flag.Int("port", 9876, "the port to send to")
	packetCount = flag.Int("count", 20, "the number of packets to send")
)

func main() {
	flag.Parse()

	hostStr := fmt.Sprintf("%s:%v", *host, *port)

	con, err := net.Dial("tcp", hostStr)
	if err != nil {
		errors.Fatal(err, "Error connecting to %s", hostStr)
	}

	b := make([]byte, 16, 16)

	var count int

	var totalTime time.Duration

	for i := 1; i <= *packetCount; i++ {
		packetStr := fmt.Sprintf("%v", i)

		start := time.Now()

		_, err = con.Write([]byte(packetStr))
		if err != nil {
			errors.Fatal(err, "Error writing to socket")
		}
		count, err = con.Read(b)

		end := time.Now()

		if err != nil {
			errors.Fatal(err, "Error reading from socket")
		}
		bStr := string(b[0:count])
		if packetStr != bStr {
			errors.Fatal(fmt.Errorf("expected %s, got %s", packetStr, bStr), "Error, invalid response")
		}

		dur := end.Sub(start)
		totalTime += dur

		fmt.Printf("%v - %s\n", i, dur.String())
	}

	avg := totalTime / time.Duration(*packetCount)

	fmt.Printf("\n%v packets sent.  Avg latency %s\n", *packetCount, avg.String())

	con.Close()
}
