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

type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() { // defaultSleeper é um Sleeper
	time.Sleep(1 * time.Second) // e o sleep dele são as pausas de 1s
}

type SpyCountdownOperations struct {
	Calls []string // ["write", "sleep", "write", "sleep"...]
}

// Sleeper! No teste a gente chama esse, que não dorme os intervalos de 1s
// ele só escreve "sleep" no log Calls []string
func (s *SpyCountdownOperations) Sleep() {
	s.Calls = append(s.Calls, sleep)
}

// escreve "write" no *SpyCountdownOperations.Calls []string
// acho que implementar interface Write() tb garante que podemos
// usar &spyCountdownOperations{} como um io.Writer
func (s *SpyCountdownOperations) Write(p []byte) (n int, err error) {
	s.Calls = append(s.Calls, write)
	return
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

// se chamar Countdown(os.Stdout, sleeper), vai escrever pra std output
// se chamar Countdown(&spyCountdownOperations, &spyCountdownOperations)
// os Fprint vao escrever usando o método Write() do spy!
// ou seja, só vão escrever "write" nas chamadas pra Fprint.
func Countdown(out io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		fmt.Fprintln(out, i)
		sleeper.Sleep()
	}
	fmt.Fprint(out, "Go!")

}

func main() {
	sleeper := &ConfigurableSleeper{1 * time.Second, time.Sleep} // passando esse sleeper, essa estrutura vai ser modificada
	Countdown(os.Stdout, sleeper)                                // vai escrever no std output mesmo, no terminal.
}
