package main

import "fmt"

// Forma 1 de declaração
func diaDaSemana1(numero int) string {
	switch numero {
	case 1:
		return "Domingo"
	case 2:
		return "Segunda-feira"
	case 3:
		return "Terça-feira"
	case 4:
		return "Quarta-feira"
	case 5:
		return "Quinta-feira"
	case 6:
		return "Sexta-feira"
	case 7:
		return "Sabado"
	default:
		return "Número inválido"
	}
}

// Forma 2 de declaração
func diaDaSemana2(numero int) string {
	switch {
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda-feira"
	case numero == 3:
		return "Terça-feira"
	case numero == 4:
		return "Quarta-feira"
	case numero == 5:
		return "Quinta-feira"
	case numero == 6:
		return "Sexta-feira"
	case numero == 7:
		return "Sabado"
	default:
		return "Número inválido"
	}
}

func main() {
	dia1 := diaDaSemana1(3)
	fmt.Println(dia1)

	dia2 := diaDaSemana2(3)
	fmt.Println(dia2)
}
