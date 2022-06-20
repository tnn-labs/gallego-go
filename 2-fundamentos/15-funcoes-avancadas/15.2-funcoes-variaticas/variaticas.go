package main

import "fmt"

// exemplo 1
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}

// exemplo 2
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5)
	fmt.Println(totalDaSoma)

	escrever("Esse é um exemplo", 10, 15, 25, 35)
}
