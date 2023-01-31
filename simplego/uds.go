package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

func runHTTPOnUDS() {
	l, err := net.Listen("unix", "/tmp/http.sock")
	defer os.Remove("/tmp/http.sock")
	// l, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}
	defer l.Close()
	m := http.NewServeMux()
	m.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("hello!, this is a HTTP server!"))
	})
	m.HandleFunc("/test", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("tested!"))
	})
	server := http.Server{
		Handler: m,
	}
	if err := server.Serve(l); err != nil {
		log.Fatal(err)
	}
}

func main() {
	// go runHTTPOnUDS()
	// Create a Unix domain socket and listen for incoming connections.
	socket, err := net.Listen("unix", "/tmp/echo.sock")
	if err != nil {
		log.Fatal(err)
	}

	// Cleanup the sockfile.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		os.Remove("/tmp/echo.sock")
		if err := os.Remove("/tmp/http.sock"); err != nil {
			log.Println(err)
		}
		os.Exit(1)
	}()

	for {
		// Accept an incoming connection.
		conn, err := socket.Accept()
		if err != nil {
			log.Fatal(err)
		}

		// Handle the connection in a separate goroutine.
		go func(conn net.Conn) {
			defer conn.Close()
			// Create a buffer for incoming data.
			buf := make([]byte, 4096)

			// Read data from the connection.
			n, err := conn.Read(buf)
			if err != nil {
				log.Fatal(err)
			}
			log.Println("get req from client:", string(buf[:n]))

			// Echo the data back to the connection.
			_, err = conn.Write(buf[:n])
			if err != nil {
				log.Fatal(err)
			}
		}(conn)
	}
}
