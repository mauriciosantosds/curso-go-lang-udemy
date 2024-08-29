package main

import "fmt"

func main() {
	//aritmeticos
	soma := 1 + 2
	subtracao := 1 -2
	var divisao float64 = 10 / 4
	multiplicacao := 10 * 5
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)
	// nao e possivel somar duas variaves com tipos diferentes exemplos int16 e int32
	var numero1 int16 = 10
	var numero2 int16 = 25
	soma2 := numero1 + numero2
	fmt.Println(soma2)
	//atribuicao
	var variavel1 string= "string"
	variavel2 := "String2"
	fmt.Println(variavel1, variavel2)
	// fim dos operadores de atribuicao

	// operadores relacionais

	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 > 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)
	// fim dos relacionais

	// operadores logicos
	fmt.Println("---------------")
	verdadeiro, falso := true, false
	fmt.Println(verdadeiro && falso)
	fmt.Println(verdadeiro || falso)
	fmt.Println(!verdadeiro)
	//fim dos operadores logicos

	//operadores unarios
	numero := 10
	numero++
	numero += 15

	numero--
	numero-=20

	numero *= 3
	numero /= 10
	numero %= 3
	fmt.Println(numero)

	var texto string
	if numero > 5 {
		texto = "Maior que 5"
	} else {
		texto = "Menor que 5"
	}

	fmt.Println(texto)

}