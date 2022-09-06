package iteration

import (
	"fmt"
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

func ExampleRepeat() {
	example := Repeat("w", 7)
	fmt.Println(example)
	//Output: wwwwwww
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 13)
	}
}
