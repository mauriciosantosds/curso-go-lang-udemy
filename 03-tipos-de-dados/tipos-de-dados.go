package main

import (
	"errors"
	"fmt"
)

func main() {
	// pega int de acordo com a arquitetura do pc se 64 bits pega int64
	numero := 10000000
	var numero2 int = 2333

	// unsygned int int sem sinal, não deixa usar int negativo (ou seja com sinal)
	var numero3 uint32 = 100
	fmt.Println(numero, numero2, numero3)

	//alias
	// alias int32 = rune
	var numero4 rune = 12313
	fmt.Println(numero4)
	//byte = uint 8
	var numero5 byte = 123

	fmt.Println(numero5)

	var numeroReal1 float32 = 123.45
	fmt.Println(numeroReal1)
	var numeroReal2 float64 = 1230000000.45
	fmt.Println(numeroReal2)
	numeroReal3 := 12345.67
	fmt.Println(numeroReal3)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto2"
	fmt.Println(str2)

	//pega numero da tabela asc do caractere
	char := 'B'
	fmt.Println(char)
    //valor 0 no go
	// string é string em branco
	//numeros é 0
	// erro é nil
	// bool é false
	var texto int
	fmt.Println(texto)

	var booleano1 bool = true
	fmt.Println(booleano1)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)
}