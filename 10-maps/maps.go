package main

import "fmt"

func main() {
	fmt.Println("Maps")
	//dentro do colchetes e o tipo das chaves e fora e o tipo dos valores
	usuario := map[string]string {
		"nome": "Pedro",
		"sobrenome": "Silva",
	}
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string {
		"nome": {
			"primeiro": "João",
			"ultimo": "Pedro",
		},
		"curso": {
			"nome": "Engenharia",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	fmt.Println("Último nome", usuario2["nome"]["ultimo"])
	delete(usuario2, "nome")
	fmt.Println(usuario2)
	usuario2["signo"] = map[string]string {
		"nome": "Gêmeos",
	}

	fmt.Println(usuario2)
}