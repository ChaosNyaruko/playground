package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

type proxy struct {
	rp *httputil.ReverseProxy
}

func (p *proxy) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Printf("serve http: %v", r.URL.String())
	p.rp.ServeHTTP(w, r)
}

func main() {
	url, err := url.Parse("https://ldoceonline.com/")
	if err != nil {
		log.Fatal(err)
	}
	p := &proxy{rp: httputil.NewSingleHostReverseProxy(url)}
	log.Fatal(http.ListenAndServe("localhost:3000", p))
}
