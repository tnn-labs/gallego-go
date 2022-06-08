package main

import (
	"fmt"
	"reflect"
)

func main() {
	// ARRAY
	var cidades [3]string
	cidades[0] = "Rio de Janeiro"
	cidades[1] = "São Paulo"
	cidades[2] = "Belo Horizonte"
	fmt.Println(cidades)

	times := [4]string{"Vasco", "Flamengo", "Fluminense", "Botafogo"}
	fmt.Println((times))

	numeros := [...]int{1, 2, 3}
	fmt.Println(numeros)

	// SLICE
	cores := []string{"azul", "vermelho", "amarelo"}
	fmt.Println(cores)

	// VERIFICANDO O TIPO
	fmt.Println(reflect.TypeOf((times)))
	fmt.Println(reflect.TypeOf((cores)))

	cores = append(cores, "verde")
	fmt.Println(cores)

	coresEscolhidas := cores[0:2]
	fmt.Println(coresEscolhidas)

	// ARRAYS INTERNOS
	// criando um array de 11 posições e pegando as 10 primeiras
	sliceExemplo := make([]float32, 10, 11)
	fmt.Println(sliceExemplo)
	fmt.Println(len(sliceExemplo)) // lenght
	fmt.Println(cap(sliceExemplo)) // capacidade

	sliceExemplo = append(sliceExemplo, 5)
	fmt.Println(sliceExemplo)
	sliceExemplo = append(sliceExemplo)
	fmt.Println(len(sliceExemplo)) // lenght
	fmt.Println(cap(sliceExemplo)) // capacidade

	sliceExemplo = append(sliceExemplo, 6)
	fmt.Println(sliceExemplo)
	sliceExemplo = append(sliceExemplo)
	fmt.Println(len(sliceExemplo)) // lenght
	fmt.Println(cap(sliceExemplo)) // capacidade
}
