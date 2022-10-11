package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool) // vai guardar o status dos websites
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			// send result to result channel
			resultChannel <- result{u, wc(u)}
		}(url) // isso aqui significa que a func anonima vai ser executada quando declarada (agora)
		// e que ela é chamada com o valor atual de 'url', o valor de cada iteracao
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}

	return results
}
