package iteration

import "testing"

func TestRepeat(t *testing.T) {
	repeated := Repeat("a")
	expected := "aaaaa"

	if repeated != expected {
		t.Errorf("Vacilão. Era pra retornar %q, mas recebemos %q", expected, repeated)
	}
}
