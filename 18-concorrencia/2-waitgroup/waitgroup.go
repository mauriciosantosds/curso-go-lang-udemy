package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//waitgroup serve para que as funcoes terminem de executar antes do programa terminar ele sincroniza as funcoes
	var waitGroup sync.WaitGroup
	waitGroup.Add(2) //qtd de go routines

	go func() {
		escrever("Olá Mundo!")
		waitGroup.Done() // -1 tira um do contador
	}()

	go func() {
		escrever("Programando em Go!")
		waitGroup.Done() // -1
	}()
	waitGroup.Wait() //pede para funcao main esperar as go routines serem executadas
}

func escrever(texto string) {
	for i := 0; i < 5; i++ {
		fmt.Println(texto)
		time.Sleep(time.Second)
	}
}
