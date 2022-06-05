package main

import "fmt"

func main() {
	// aritmeticos
	// -----------
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	// atribuição
	// ----------
	var variavel1 string = "String 1"
	variavel2 := "String 2"
	fmt.Println(variavel1, variavel2)

	// relacionais
	// -----------
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// lógicos
	// -------
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	fmt.Println(!falso)

	// unários
	// -------
	numero := 10
	numero++     // numero = numero + 1
	numero += 15 // numero = numero + 15
	fmt.Println(numero)

	numero--     // numero = numero - 1
	numero -= 20 // numero = numero - 20
	fmt.Println(numero)

	numero *= 3 // numero = numero * 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	// ternario -> não existe -> utilizar if else
	// --------
	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "menor que 5"
	}
	fmt.Println(texto)
}
