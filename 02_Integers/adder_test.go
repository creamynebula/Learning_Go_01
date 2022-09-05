package integers

import (
	"fmt"
	"testing"
)

func TestAdderino(t *testing.T) { //veja q desde q o nome começa com TestAdd, ele reconhece a funcao Add
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Esperávamos %d, mas recebemos %d. GRRRRRRR.", expected, sum) //%d imprime int. %q imprimia string.
	}

}

func ExampleAdd() {
	//se o comentário embaixo tivesse 7 o valor, ia dar fail o teste! essa funcao executa teste automatico igual a outra!
	sum := Add(1, 5)
	fmt.Println(sum)
	// Output: 6

}
