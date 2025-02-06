package main

import (
	"fmt"
	"io"
	"os"
)

const finalWorld = "Go!"
const countdonwStart = 3

func Countdown(w io.Writer) {
	for i := countdonwStart; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprint(w, finalWorld)
}

func main() {
	Countdown(os.Stdout)
}
