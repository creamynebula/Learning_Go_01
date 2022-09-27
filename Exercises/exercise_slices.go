package main

/*
import (
	"golang.org/x/tour/pic"
)
na versão web do tour isso funciona, aqui nao, entao vou comentar isso e o main pra nao receber o warning
*/

func Pic(dx, dy int) [][]uint8 {

	s1 := make([][]uint8, dy)

	for y := 0; y < dy; y++ {
		s1[y] = make([]uint8, dx)

		for x := 0; x < dx; x++ {
			s1[y][x] = uint8((x + y) / 2)
		}
	} // fim Pic

	return s1
}

func main() {
	//pic.Show(Pic)
}
