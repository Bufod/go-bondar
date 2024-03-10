package main

import (
	"github.com/Bufod/go-bondar/cmd/shortener/handlers"
)

func main() {
	router := handlers.SetupRouter()
	router.Run()
}
