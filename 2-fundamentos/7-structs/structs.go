package main

import "fmt"

type usuario struct {
	nome     string
	idade    uint8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     uint8
}

func main() {
	// exemplo 1
	var user1 usuario
	user1.nome = "Jhon"
	user1.idade = 21
	fmt.Println(user1)

	// exemplo 2 -> com inferencia de tipo
	enderecoExemplo := endereco{"Rua Ali perto", 0}
	user2 := usuario{"Mary", 25, enderecoExemplo}
	fmt.Println(user2)

	// exemplo 3 -> omitindo alguma prop
	user3 := usuario{nome: "bob"}
	fmt.Println(user3)
}
