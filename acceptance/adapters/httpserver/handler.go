package httpserver

import (
	"fmt"
	"net/http"

	greet "valentine.com/tdd/acceptance/domain"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	name := r.URL.Query().Get("name")
	fmt.Fprintf(w, "%s", greet.Greet(name))
}
