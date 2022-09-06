package iteration

import (
	"strings"
	"testing"
)

const numRepeatings = 7

func TestRepeat(t *testing.T) {
	got := Repeat("a", numRepeatings)
	expected := strings.Repeat("a", numRepeatings)

	if got != expected {
		t.Errorf("Vacilão. Era pra retornar %q, mas recebemos %q.", expected, got)
	}
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 13)
	}
}
