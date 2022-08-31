package main //a package is a way to group functions. we declared a package called "main"
//the package is made up of all the files in the same directory

import (
	"fmt"

	"rsc.io/quote"
)

//imports the 'fmt' package, which cointains functions for formatting text. part of the standard library.

func Hello() string {
	return "Oláaaaaa, enfermeira!!\n"
}

func main() {
	fmt.Println(Hello())
	fmt.Println(quote.Go())
} //main function implemented
