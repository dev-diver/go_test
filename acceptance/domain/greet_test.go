package greet_test

import (
	"testing"

	greet "valentine.com/tdd/acceptance/domain"
	specifications "valentine.com/tdd/acceptance/specification"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(greet.Greet))
}
