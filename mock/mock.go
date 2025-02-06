package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWorld = "Go!"
const countdonwStart = 3

type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() {
	s.Calls++
}

func Countdown(w io.Writer) {
	for i := countdonwStart; i > 0; i-- {
		fmt.Fprintln(w, i)
		time.Sleep(1 * time.Second)
	}
	fmt.Fprint(w, finalWorld)
}

func main() {
	Countdown(os.Stdout)
}
