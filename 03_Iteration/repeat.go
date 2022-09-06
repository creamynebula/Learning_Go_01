package iteration

func Repeat(character string, numRepeatings int) string {

	var repeatCount = numRepeatings //numero de vezes que vamos repetir a letra
	var repeated string             //em Go, repeated == ""

	for i := 0; i < repeatCount; i++ {
		repeated += character
	}

	return repeated
}
