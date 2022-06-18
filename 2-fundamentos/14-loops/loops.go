package main

import (
	"fmt"
)

func main() {
	// EXEMPLO 1
	// ---------
	// i := 0
	// for i < 10 {
	// 	i++
	// 	fmt.Println("Implementando i")
	// 	time.Sleep(time.Second)
	// }
	// fmt.Println(i)

	// EXEMPLO 2
	// ---------
	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("Implementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// EXEMPLO 3
	// ---------
	// nomes := [3]string{"Joao", "Davi", "Lucas"}
	// for indice, nome := range nomes {
	// 	fmt.Println(indice, nome)
	// }

	// EXEMPLO 4
	// ---------
	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice, string(letra))
	// }

	// EXEMPLO 5
	// ---------
	usuario := map[string]string{
		"nome":      "Jhon",
		"sobrenome": "Doe",
	}

	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}
}
