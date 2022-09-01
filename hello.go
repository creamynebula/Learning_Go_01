package main //a package is a way to group functions. we declared a package called "main"
//the package is made up of all the files in the same directory

import (
	"fmt"

	"rsc.io/quote"
) //imports the 'fmt' package, which cointains functions for formatting text. part of the standard library.

const prefixHello = "Olá, "
const prefixHola = "Hola, "

func Hello(name string, language string) string {

	var prefix = prefixHello

	if name == "" {
		name = "enfermeira"
	}
	if language == "Spanish" {
		prefix = prefixHola
	}

	return prefix + name + "!"
}

func Chihaya() string {
	chihaya := "\nsincerely,\n\tAyase Chihaya.\n" //usually I would be using a constant, obviously.
	return chihaya
}

func main() {
	fmt.Println(Hello("Oe Kanade-chan", ""))
	fmt.Println(quote.Go())
	fmt.Println(Chihaya())
} //main function implemented
