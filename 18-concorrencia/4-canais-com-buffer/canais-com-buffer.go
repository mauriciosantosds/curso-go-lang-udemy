package main

import "fmt"

func main() {
	// canal com buffer especifica capacidade 2, so vai bloquear se atingir a capacidade maxima
	canal := make(chan string, 2)

	canal <- "Olá Mundo!"
	canal <- "Programando em go"
	//canal <- "Terceiro valor!" causa deadlock

	mensagem := <-canal
	mensagem2 := <-canal
	mensagem3 := <-canal
	fmt.Println(mensagem)
	fmt.Println(mensagem2)
	fmt.Println(mensagem3) //causa deadlock
}
