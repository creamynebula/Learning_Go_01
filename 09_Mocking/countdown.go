package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const countdownStart = 3
const write = "write"
const sleep = "sleep"

// qualquer tipo que implemente Sleep() é um Sleeper.
type Sleeper interface {
	Sleep()
}

/*
type SpySleeper struct {
	Calls int
}

func (s *SpySleeper) Sleep() { // então, SpySleeper é um Sleep
	s.Calls++
}
*/

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() { // defaultSleeper é um Sleep
	time.Sleep(1 * time.Second) // e o sleep dele são as pausas de 1s
}

type SpyCountdownOperations struct {
	Calls []string // ["write", "sleep", "write", "sleep"...]
}

// Sleeper! No teste a gente chama esse, pq não tem os intervalos de 1s
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// escreve "write" no slice de calls ao spy
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
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
