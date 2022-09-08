package main

func Sum(numbers [5]int) int {
	var result = 0
	for _, elemento := range numbers {
		result += elemento
	}

	return result
}
