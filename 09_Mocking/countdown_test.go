package main

import (
	"bytes"
	"testing"
)

func TestCountdown(t *testing.T) {
	// Buffer é um slice de bytes, que vc pode ler e escrever
	// buffer = &bytes.Buffer pq daí vc vai mexer nesse slice
	buffer := &bytes.Buffer{}
	spySleeper := &SpySleeper{} // sleeper com Calls == 0

	Countdown(buffer, spySleeper)

	got := buffer.String()
	want := `3
2
1
Go!`

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}

	if spySleeper.Calls != 3 {
		t.Errorf("not enough calls to sleeper. want 3, got %d", spySleeper.Calls)
	}
}
