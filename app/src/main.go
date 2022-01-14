package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("HelloWorld"))
	})

	if err := http.ListenAndServe(":5045", nil); err != nil {
		log.Fatal(err)
	}
}
