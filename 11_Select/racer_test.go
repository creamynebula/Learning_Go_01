package main

import (
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

func TestRacer(t *testing.T) {

	t.Run("compares speeds of servers, returning the url of the fastest one", func(t *testing.T) {
		// o argumento de makeDelayedServer é o delay pro tempo de resposta
		slowServer := makeDelayedServer(20 * time.Millisecond)
		fastServer := makeDelayedServer(0 * time.Millisecond)
		// defer significa que eles vão ser executados
		// no fim da funcao que os contém (TestRacer)
		defer slowServer.Close()
		defer fastServer.Close()

		slowURL := slowServer.URL
		fastURL := fastServer.URL
		// Racer(url1, url2) retorna a url do site que abrir 1o
		want := fastURL
		// só retorna erro se der 'timeout', default == 10s
		got, err := Racer(slowURL, fastURL)

		if err != nil {
			t.Fatalf("did not expect an error but got one %v", err)
		}

		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	})

	t.Run("returns an error if a server doesn't respond within the specified time", func(t *testing.T) {
		// server com delay de 25ms
		server := makeDelayedServer(25 * time.Millisecond)
		defer server.Close()

		// só que o timeout do nosso racer é 20ms, então tem que ter erro
		_, err := ConfigurableRacer(server.URL, server.URL, 20*time.Millisecond)

		if err == nil {
			t.Error("expected an error but didn't get one")
		}
	})
}

func makeDelayedServer(delay time.Duration) *httptest.Server {
	return httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// nosso server dorme 'delay', e dps dá bom
		time.Sleep(delay)
		w.WriteHeader(http.StatusOK)
	}))
}
