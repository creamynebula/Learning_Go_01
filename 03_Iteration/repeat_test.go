package iteration

import (
	"strings"
	"testing"
)

func TestRepeat(t *testing.T) {
	numRepeatings := 7
	repeated := Repeat("a", numRepeatings)
	expected := strings.Repeat("a", numRepeatings)

	if repeated != expected {
		t.Errorf("Vacilão. Era pra retornar %q, mas recebemos %q", expected, repeated)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 13)
	}
}
