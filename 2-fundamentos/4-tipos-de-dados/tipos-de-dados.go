package main

import (
	"errors"
	"fmt"
)

func main() {
	// tipos de numeros inteiros
	// -------------------------
	// int8, int16, int32, int64
	// int -> utiliza a arquitetura do computador como base

	var numero1 int8 = 100

	// neste exemplo está sendo inferido o tipo int com a quantidade de bits do computador local
	numero2 := 1000000000

	fmt.Println(numero1)
	fmt.Println(numero2)

	// uint -> unsygned int -> int sem sinal
	var numero3 uint32 = 10000
	fmt.Println(numero3)

	// alias
	// int32 = rune
	var numero4 rune = 123456
	fmt.Println(numero4)

	// byte = uint8
	var numero5 byte = 123
	fmt.Println(numero5)

	// numeros reais
	// -------------
	// float32, float64

	var numeroReal1 float32 = 1023890003.05
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 1023894893029200202930393039303030330003.45
	fmt.Println(numeroReal2)

	numeroReal3 := 1234.78
	fmt.Println(numeroReal3)

	// strings
	// -------
	var str1 string = "Texto um"
	fmt.Println(str1)

	str3 := "Texto dois"
	fmt.Println(str3)

	char := 'B'
	fmt.Println(char)

	// valor zero -> valor inicial quando nenhum valor é atribuido
	// ------------
	var texto1 int
	fmt.Println(texto1)

	var texto2 string
	fmt.Println(texto2)

	// tipo booleano
	// seu valor inicial é false
	var booleano1 bool = true
	fmt.Println(booleano1)

	// tipo especifico para erros
	// valor inicial é <nil>
	var erro1 error
	fmt.Println(erro1)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
