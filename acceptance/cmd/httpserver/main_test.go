package main_test

import (
	"testing"

	specifications "valentine.com/tdd/acceptance/specification"
)

func TestHTTPServer(t *testing.T) {
	specifications.GreetSpecification(t, &HTTPServer{})
}
