package main

import (
	"amazon/lib/handler"
	"fmt"
	"net/http"
)

var (
	port = ":6000"
)

func main() {
	fmt.Println("starting server...")

	go handler.Main()

	http.ListenAndServe(port, nil)
}
