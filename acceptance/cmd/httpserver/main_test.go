package main_test

import (
	"testing"

	go_specs_greet "valentine.com/tdd/acceptance"
	specifications "valentine.com/tdd/acceptance/specification"
)

func TestHTTPServer(t *testing.T) {
	driver := go_specs_greet.Driver{BaseURL: "http://localhost:8080"}
	specifications.GreetSpecification(t, driver)
}
