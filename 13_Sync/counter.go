package main

import "sync"

type Counter struct {
	value int
	mu    sync.Mutex // 'zero value' == destrancado
}

func (c *Counter) Inc() {
	// primeiro trancar. aí se outra goroutine chamar esse método,
	// ela fica blockada, até nós destrancarmos.
	c.mu.Lock()         // porta trancada
	defer c.mu.Unlock() // destrancada
	c.value += 1
}

func (c *Counter) Value() int {
	return c.value
}

func main() {

}
