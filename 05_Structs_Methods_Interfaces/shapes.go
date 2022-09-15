package main

import "math"

type Rectangle struct { // retângulos tem altura e largura
	Width  float64
	Height float64
}

/* método que calcula a área de um retângulo
ele age sobre (rectangle Rectangle), é chamado fazendo rectangle.Area(), retorna um float64 */
func (rectangle Rectangle) Area() float64 { // metodo
	return (rectangle.Width * rectangle.Height)
}

type Circle struct {
	Radius float64
}

func (circle Circle) Area() float64 {
	return (circle.Radius * circle.Radius * math.Pi)
}

type Shape interface {
	Area() float64
} //isso quer dizer que Shape é um struct que possui um método Area()
// nós realmente implementamos esse método pra Circle e Rectangle, que são nossos
// atuais usos de Shape

func Perimeter(rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}
