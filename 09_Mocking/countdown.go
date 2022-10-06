package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3

// qualquer tipo que implemente Sleep() é um Sleeper.
type Sleeper interface {
	Sleep()
}

type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() { // então, SpySleeper é um Sleep
	s.Calls++
}

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() { // defaultSleeper tb é um Sleep
	time.Sleep(1 * time.Second)
}

func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "Go!")

}

func main() {
	sleeper := &DefaultSleeper{}
	Countdown(os.Stdout, sleeper)
}
