package main

import (
	"net/http"
)

func main() {
	handler := http.HandlerFunc(Handler)
	http.ListenAndServe(":8080", handler)
}
