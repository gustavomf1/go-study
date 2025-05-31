package main

import (
	"errors"
	"fmt"
)

func main() {
	var numero int16 = 10000
	fmt.Println(numero)

	nume2 := 1000000000
	fmt.Println(nume2)

	// uint nao pode ser negativo
	var nume3 uint32 = 100
	fmt.Println(nume3)

	var numeroReal1 float32 = 123.55
	var numeroReal2 float64 = 12342424242424.345

	fmt.Println(numeroReal1, numeroReal2)

	var str string = "Texto"
	fmt.Println(str)

	str2 := "Texto 2"
	fmt.Println(str2)

	char := 'B'
	fmt.Println(char)

	var texto int
	fmt.Println(texto)

	var booleanol bool = true

	fmt.Println(booleanol)

	var erro error
	fmt.Println(erro)

	var erro2 error = errors.New("Erro interno")
	fmt.Println(erro2)
}
