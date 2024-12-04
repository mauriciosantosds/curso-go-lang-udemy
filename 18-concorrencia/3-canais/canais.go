package main

import (
	"fmt"
	"time"
)

// o clase é responsável por fechar o canal e depois é verificado no loop se esse canal ainda
// está aberto o significa que pode receber dados caso contrário o programa é encerrado
func main() {
	canal := make(chan string)
	go escrever("Olá Mundo!", canal)
	fmt.Println("Depois da função escrever começar a ser executada")
	/* for {
		mensagem, aberto := <-canal //espera canal
		if !aberto {
			break
		}
		fmt.Println(mensagem)
	} */

	for mensagem := range canal {
		fmt.Println(mensagem)
	}

	fmt.Println("Fim do programa!")
}

func escrever(texto string, canal chan string) {
	//time.Sleep(time.Second * 5)
	for i := 0; i < 5; i++ {
		canal <- texto //passa valor para o canal
		time.Sleep(time.Second)
	}
	close(canal)
}
