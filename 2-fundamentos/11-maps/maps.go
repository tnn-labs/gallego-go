package main

import "fmt"

func main() {
	usuario1 := map[string]string{
		"nome":      "Jhon",
		"sobrenome": "Doe",
	}

	fmt.Println(usuario1)
	fmt.Println(usuario1["nome"])
	fmt.Println(usuario1["sobrenome"])

	// map aninhado
	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro":  "Jane",
			"sobrenome": "Doe",
		},
		"curso": {
			"nome":   "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)

	// deletando chave
	delete(usuario2, "campus")
	fmt.Println(usuario2)

	// adicionar chave
	usuario2["Natural"] = map[string]string{
		"cidade": "Rio de Janeiro",
	}
	fmt.Println(usuario2)
}
