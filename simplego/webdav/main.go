package main

import (
	"log"
	"net/http"

	"golang.org/x/net/webdav"
)

func main() {
	h := &webdav.Handler{
		Prefix:     "/dav/",
		FileSystem: webdav.Dir("."),
		LockSystem: webdav.NewMemLS(),
		Logger: func(r *http.Request, err error) {
			log.Printf("req: %+v, err: %v", r, err)
		},
	}

	http.Handle("/dav/", h)

	if err := http.ListenAndServe(":9091", nil); err != nil {
		log.Fatalf("%v", err)
	}
}
