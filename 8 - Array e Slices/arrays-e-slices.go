package main

import (
	"fmt"
	"reflect"
)

func main() {
	var array1 [4]string
	array1[0] = "Gustavo"
	array1[1] = "Martins"
	array1[2] = "França"
	array1[3] = "Magnossão"
	fmt.Println(array1)

	array2 := [4]string{"Magnossão", "França", "Martins", "Gustavo"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice := []int{1, 2, 3}
	fmt.Println(slice)

	slice = append(slice, 22)
	fmt.Println(slice)

	fmt.Println(reflect.TypeOf(slice))
	fmt.Println(reflect.TypeOf(array1))

	slice2 := array2[0:2]
	fmt.Println(slice2)
}
