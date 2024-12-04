package main

import (
	"fmt"
	"time"
)

func main() {
	go escrever("Olá mundo!")      //goroutine não espera acabar para seguir para a próxima linha
	escrever("Programando em Go!") // se colocar go em tudo ele não vai imprimir nada pq não espera executar e termina o programa
}

func escrever(texto string) {
	for i := 0; i < 2; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
