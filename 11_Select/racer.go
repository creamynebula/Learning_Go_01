package main

import (
	"fmt"
	"net/http"
	"time"
)

// por default, é pros requests darem timeout se nao tiver resposta em 10s
var tenSecondTimeout = 10 * time.Second

func Racer(a, b string) (winner string, error error) {
	return ConfigurableRacer(a, b, tenSecondTimeout)
}

// pra gente poder testar racer com um timeout menor que o default
func ConfigurableRacer(a, b string, timeout time.Duration) (winner string, error error) {
	select {
	case <-ping(a): // se conseguir abrir a url 'a', retorna ela
		return a, nil
	case <-ping(b):
		return b, nil
	case <-time.After(timeout): // se demorar mais que timeout
		return "", fmt.Errorf("timed out waiting for %s and %s", a, b)
	}
}

func ping(url string) chan struct{} {
	// função retorna um channel struct vazio, um pointer,
	// pq isso não ocupa memória. a idéia é retornar algo
	// que nao gasta recurso assim que tiver aberto a url,
	// com o objetivo de sinalizar que deu bom abrir a url apenas.
	ch := make(chan struct{})
	go func() {
		http.Get(url) // isso blocka
		close(ch)     // isso tb blocka
	}()
	return ch // isso blocka, pq depende de close(ch), então só retorna se tiver aberto url
}
