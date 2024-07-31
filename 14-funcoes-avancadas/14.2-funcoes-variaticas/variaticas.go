package main

import "fmt"

func soma(numeros ...int) int {
	total := 0

	for _, numero := range numeros {
		total += numero
	}

	return total
}
//não pode ter mais de um parâmetro variático por função e tem que ser o último
func escrever(texto string, numeros ...int) {
	for _, numero := range numeros {
		fmt.Println(texto, numero)
	}
}

func main() {
	totalDaSoma := soma(1, 2, 3, 4, 5, 6)
	totalDaSomaVazio1 := soma()
	fmt.Println(totalDaSoma)
	fmt.Println(totalDaSomaVazio1)

	escrever("Olá mundo", 10, 2, 3, 4)
}