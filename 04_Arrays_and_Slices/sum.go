package main

import "fmt"

var numbers2 = [...]int{0, 0, 0, 0, 28}

//recebe um vetor e retorna a soma dos elementos
func Sum(numbers []int) int {
	var result = 0
	for _, elemento := range numbers {
		result += elemento
	}

	return result
}

//recebe uma lista de vetores, retorna um vetor contendo as somas dos elementos dos respectivos vetores originais
//ex: SumAll({1,2}, {3,4}) -> [3 7]
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
	fmt.Println(Sum([]int{7, 7, 7, 7}))
	fmt.Println(SumAll([]int{1, 2}, []int{3, 4}))
}
