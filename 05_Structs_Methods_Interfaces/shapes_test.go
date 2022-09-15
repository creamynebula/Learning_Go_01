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

/* table driven tests que chama isso */
func TestArea(t *testing.T) {

	areaTests := []struct { //areaTests é um slice de structs
		shape Shape   //que contém um shape
		want  float64 //e a área desse shape
	}{ //aí aqui declaramos literalmente 2 exemplos, que vamos usar nos testes
		{Rectangle{12, 6}, 72.0},
		{Circle{10}, 314.1592653589793},
	}

	for _, tt := range areaTests { //iterar sober areaTests
		got := tt.shape.Area() // got == área do shape que tá ali colocado de exemplo
		if got != tt.want {    // se nao for igual a área real do shape, calculada na mão
			t.Errorf("got %g want %g", got, tt.want)
		}
	}

}
