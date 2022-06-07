package main

import "fmt"

func main() {
	fmt.Println("----------> exemplo variavel")
	var variavel1 int = 100
	var variavel2 int = variavel1

	fmt.Println(variavel1, variavel2)

	variavel1 = 150
	fmt.Println(variavel1, variavel2)

	fmt.Println("----------> exemplo ponteiro")
	// ponteiro é uma referência de memória
	var variavel3 int
	var ponteiro *int

	variavel3 = 100
	ponteiro = &variavel3

	fmt.Println(variavel3, ponteiro, *ponteiro) // *ponteiro -> desreferenciação

	variavel3 = 150
	fmt.Println(variavel3, ponteiro, *ponteiro)
}
