package main

import (
	"github.com/Bufod/go-bondar/cmd/shortener/handlers"
	"net/http"
)

func main() {
	http.HandleFunc("/", handlers.MainHandler)
	http.ListenAndServe(":8080", nil)
}
