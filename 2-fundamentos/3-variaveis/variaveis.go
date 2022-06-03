package main

import "fmt"

func main() {
	// atribuição individual
	var variavel1 string = "Variavel 1"
	variavel2 := "Variavel 2"

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	// atribuição em grupo
	var (
		variavel3 string = "Variavel 3"
		variavel4 string = "Variavel 4"
	)

	fmt.Println(variavel3)
	fmt.Println(variavel4)

	// atribuição em grupo com inferencia de tipo
	variavel5, variavel6 := "Variavel 5", "Variavel 6"
	fmt.Println(variavel5)
	fmt.Println(variavel6)
}
