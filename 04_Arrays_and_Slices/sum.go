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

/* recebe uma lista de vetores, retorna um vetor contendo as somas dos elementos
dos respectivos vetores originais.
ex: SumAll([]int{1,2}, []int{3,4}) -> [3 7] */
func SumAll(vectorsToSum ...[]int) []int {

	numberOfVectors := len(vectorsToSum)
	result := make([]int, numberOfVectors) //aloca um vetor de tamanho 'numberOfVectors'

	for i, vector := range vectorsToSum {
		result[i] = Sum(vector)
	}

	return result
}

/* mesma coisa, mas, porque acessar result[i] quando ele nao existe dá um runtime error,
fazer result = append(result, novoElemento) é mais seguro. */
func SumAll2(vectorsToSum ...[]int) []int {
	var result []int

	for _, vector := range vectorsToSum {
		result = append(result, Sum(vector)) //append(vetor, valor) retorna |vetor valor|
	}

	return result
}

/* retorna as somas das caudas de uma lista de vetores. cauda é todo o vetor
menos o 1o elemento. */
func SumAllTails(numbersToSum ...[]int) []int {

	var sums []int // vai segurar a soma das caudas

	for _, numbers := range numbersToSum { //pra cada vetor "numbers" na coleção
		tail := numbers[1:]            //tail = novo vetor, cópia a partir do elemento índice 1 de 'numbers'
		sums = append(sums, Sum(tail)) //coloque em sums a soma da cauda do vetor atual
	}

	return sums
}

func main() {
	fmt.Println(numbers2)
	fmt.Println(Sum([]int{7, 7, 7, 7}))
	fmt.Println(SumAll([]int{1, 2}, []int{3, 4}))
	fmt.Println(SumAll2([]int{1, 2}, []int{3, 4}))
}
