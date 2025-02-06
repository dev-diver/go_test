package main

import (
	"fmt"
	"io"
	"os"
)

func Countdown(w io.Writer) {
	fmt.Fprintf(w, "3")
}

func main() {
	Countdown(os.Stdout)
}
