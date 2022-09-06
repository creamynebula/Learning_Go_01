package iteration

func Repeat(character string, numRepeatings int) string {

	var repeated string //em Go, repeated == ""

	for i := 0; i < numRepeatings; i++ {
		repeated += character
	}

	return repeated
}
