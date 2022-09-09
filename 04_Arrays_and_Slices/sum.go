package main

import "fmt"

var numbers2 = [...]int{0, 0, 0, 0, 28}

func Sum(numbers []int) int {
	var result = 0
	for _, elemento := range numbers {
		result += elemento
	}

	return result
}

func SumAll(numbersToSum ...[]int) []int {
	numberOfVectors := len(numbersToSum)
	result := make([]int, numberOfVectors)

	for i, vector := range numbersToSum {
		result[i] = Sum(vector)
	}

	return result
}

func main() {
	fmt.Println(numbers2)
}
