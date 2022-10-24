package main

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

		assertCounter(t, &counter, 3)
	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCount := 1000
		counter := Counter{}

		var wg sync.WaitGroup // um WaitGroup espera por uma coleçao de goroutines terminar
		wg.Add(wantedCount)   // número de goroutines pelas quais teremos de esperar
		// é bom Add() antes de chamar as goroutines

		for i := 0; i < wantedCount; i++ {
			go func() { // nova goroutine
				counter.Inc() // incremente o counter
				wg.Done()     // chamar 'done' qdo terminar, me parece que isso é tipo wg.Add(-1)
			}()
		}

		wg.Wait() // espera todas as goroutines terminarem
		// ou seja, blocka até o counter interno do WaitGroup == 0

		assertCounter(t, &counter, wantedCount)
	})
}

func assertCounter(t testing.TB, got *Counter, want int) {
	t.Helper()
	if got.Value() != want {
		t.Errorf("got %d, want %d", got.Value(), want)
	}
}
