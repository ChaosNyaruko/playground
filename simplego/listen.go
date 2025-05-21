package main

import (
	"fmt"
	"net"
)

func main() {
	l, err := net.ListenPacket("udp", "127.0.0.1:9006")
	if err != nil {
		panic(err)
	}
	for {
		b := make([]byte, 1400)
		n, conn, err := l.ReadFrom(b)
		if err != nil {
			panic(err)
		}
		fmt.Printf("UDP packet: %q, conn: %s\n", string(b[:n]), conn.String())
		l.WriteTo([]byte(fmt.Sprintf("You are %s\n", conn.String())), conn)
	}
}
