package main

import (
	"fmt"
	"log"
	"net/http"
)

const (
	a = 2 << iota
	b
	c
	d
)

func main() {
	//http.HandleFunc("/foo", fooHandler)

	http.HandleFunc("/bar", func(w http.ResponseWriter, r *http.Request) {
		_, err := fmt.Fprintf(w,r.Method)
		if err != nil {
			return
		}
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func fooHandler(w http.ResponseWriter, r *http.Request) string {
	return "hello"
}
