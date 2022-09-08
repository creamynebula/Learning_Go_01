package main

func Sum(numbers [5]int) int {
	var result = 0
	for i := 0; i < 5; i++ {
		result += numbers[i]
	}

	return result
}
