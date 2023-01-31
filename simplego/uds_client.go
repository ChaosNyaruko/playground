package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

func main() {
	send()
}

func send() {
	for i := 0; i < 10000; i++ {
		conn, err := net.Dial("unix", "/tmp/echo.sock")
		if err != nil {
			log.Fatal(err)
		}

		msg := fmt.Sprintf("Hello, this is %d", i)
		if _, err := conn.Write([]byte(msg)); err != nil {
			log.Fatal(err)
		}

		b := make([]byte, len(msg))
		if _, err := conn.Read(b); err != nil {
			log.Fatal(err)
		}
		log.Println("get resp from server:", string(b))
		time.Sleep(time.Second)
	}
}
