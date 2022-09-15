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
		name    string  // o nome é pra gente usar no nome de cada teste, pra depois os testes serem mais legíveis
		shape   Shape   //que contém um shape
		hasArea float64 //e a área desse shape
	}{ //aí aqui declararemos literalmente exemplos, que vamos usar nos testes
		{name: "Rectangle", shape: Rectangle{Width: 12, Height: 6}, hasArea: 72.0},
		{name: "Circle", shape: Circle{Radius: 10}, hasArea: 314.1592653589793},
		{name: "Triangle", shape: Triangle{Base: 12, Height: 6}, hasArea: 36.0},
		/* os termos 'name', 'shape' e 'want' poderiam ter sido omitidos aqui, podia ter sido
		{"Rectangle", Rectangle{12, 6}, 72.0},
		{"Circle", Circle{10}, 314.1592653589793},
		{"Triangle", Triangle{12, 6}, 36.0},
		*/
	}

	for _, tt := range areaTests { //iterar sober areaTests
		t.Run(tt.name, func(t *testing.T) {
			got := tt.shape.Area() // got == área do shape que tá ali colocado de exemplo
			if got != tt.hasArea { // se nao for igual a área real do shape, calculada na mão
				t.Errorf("got %g want %g", got, tt.hasArea)
			}
		})
	}
}
