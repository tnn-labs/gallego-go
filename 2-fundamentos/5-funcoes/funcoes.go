package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	subtracao := n1 - n2
	return soma, subtracao
}

func main() {
	// exemplo 1
	soma := somar(10, 20)
	fmt.Println(soma)

	// exemplo 2
	var f1 = func() {
		fmt.Println("Função 1")
	}

	f1()

	// exemplo 3
	var f2 = func(txt string) string {
		fmt.Println(txt)
		return txt
	}

	resultado := f2("texto para a função f2")
	fmt.Println(resultado)

	// exemplo 4
	resultadoSoma1, resultadoSubtracao1 := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma1, resultadoSubtracao1)

	resultadoSoma2, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadoSoma2)
}
