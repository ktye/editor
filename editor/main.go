package main

//go:generate go run gen.go

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	var httpAddr string
	flag.StringVar(&httpAddr, "http", "127.0.0.1:2017", "HTTP service address, e.g. 127.0.0.1:2017")
	flag.Parse()

	http.HandleFunc("/", handler)
	if err := http.ListenAndServe(httpAddr, nil); err != nil {
		log.Fatal(err)
	}
}
