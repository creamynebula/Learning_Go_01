package main //a package is a way to group functions. we declared a package called "main"
//the package is made up of all the files in the same directory

import (
	"fmt"

	"rsc.io/quote"
) //imports the 'fmt' package, which cointains functions for formatting text. part of the standard library.

const prefixHello = "Olá, "

func Hello(name string) string {
	//return "\nOláaaaaa, enfermeira!!\n"
	return prefixHello + name + "!"
}

func Chihaya() string {
	chihaya := "\nAyase Chihaya!\n"
	return chihaya
}

func main() {
	fmt.Println(Hello("Kanade"))
	fmt.Println(quote.Go())
	fmt.Println(Chihaya())
} //main function implemented
