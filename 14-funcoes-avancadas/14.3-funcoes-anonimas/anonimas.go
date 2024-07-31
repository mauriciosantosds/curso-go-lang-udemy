package main

import "fmt"

func main() {
	retorno := func(texto string) string {
		fmt.Println(texto)
		return fmt.Sprintf("recebido -> %s %d", texto, 10)
	}("teste")
	fmt.Println(retorno)
}