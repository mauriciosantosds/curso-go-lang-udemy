package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1[5] int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5]string{"Posição1", "Posição 2", "Posição 3"}
	fmt.Println(array2)
	array2[4] = "teste"
	array3 := [...]int {1,2,3,4,5}
	fmt.Println(array3)

	slice := []int{10, 12, 13, 14, 15, 16, 17}
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array3))

	slice = append(slice, 18)

	fmt.Println(slice)

	slice2 := array2[1:3]
	fmt.Println(slice2)

	array2[1] = "Posição alterada"
	fmt.Println(slice2)
}