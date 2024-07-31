package main

import (
	"fmt"
	"time"
)

func main() {
/* 	i := 0

	for i < 10 {
		i++
		fmt.Println("Incrementando i")
		time.Sleep(time.Second)
	}

	fmt.Println(i)
// pode usar j += 2 para aumentar de 2 em 2
	for j := 0; j < 10; j++ {
		fmt.Println("Incrmentando j", j)
		time.Sleep(time.Second)
	} */

	nomes := [3]string{"João", "Davi", "Lucas"}

	for indice, nome := range nomes {
		fmt.Println(indice, nome)
	}

	for _, nome := range nomes {
		fmt.Println(nome)
	}
	//retorna codigo da tabela asc da letra
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, letra)
	}

	//retorna codigo da tabela asc da letra
	for indice, letra := range "PALAVRA" {
		fmt.Println(indice, string(letra))
	}

	usuario := map[string]string {
		"nome": "Leonardo",
		"sobrenome": "Silva",
	}

	fmt.Println(usuario)

	for chave, valor := range usuario {
		fmt.Println(chave, valor)
	}

	for {
		fmt.Println("Executando infinitamente")
		time.Sleep(time.Second)
	}

}