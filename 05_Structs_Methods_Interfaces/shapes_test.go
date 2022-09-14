package main

import (
	"testing"
)

func TestPerimeter(t *testing.T) {
	rectangle := Rectangle{10.0, 10.0}
	got := Perimeter(rectangle)
	want := 40.0

	if got != want {
		t.Errorf("recebemos %.2f, queríamos %.2f. Vacilo!", got, want)
	}

}

func TestArea(t *testing.T) {
	rectangle := Rectangle{2.0, 4.0}
	got := Area(rectangle)
	want := 8.0

	if got != want {
		t.Errorf("recebemos %.2f, queríamos %.2f. Vacilo!", got, want)
	}
}
