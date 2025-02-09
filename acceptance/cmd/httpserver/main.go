package main

import (
	"net/http"

	"valentine.com/tdd/acceptance/adapters/httpserver"
)

func main() {
	handler := http.HandlerFunc(httpserver.Handler)
	http.ListenAndServe(":8080", handler)
}
