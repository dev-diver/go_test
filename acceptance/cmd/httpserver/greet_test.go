package main_test

import (
	"testing"

	main "valentine.com/tdd/acceptance/cmd/httpserver"
	specifications "valentine.com/tdd/acceptance/specification"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(main.Greet))
}
