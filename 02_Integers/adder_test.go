package integers

import "testing"

func TestAdder(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("Esperávamos %d, mas recebemos %d. GRRRRRRR.", expected, sum) //%d imprime int. %q imprimia string.
	}

}
