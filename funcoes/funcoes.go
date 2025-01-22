package main

import "fmt"

func main() {
	fmt.Println(soma(10, 10))
	fmt.Println(sub(12.750, 2.750))
	fmt.Println(div(11.5, 2))
	fmt.Println(mult(10))
	fmt.Println(texto("Gustavo Magnusson"))
	valor, valor2 := Soma2(10, 5)
	fmt.Println(valor)
	fmt.Println(valor2)
}

func soma(x int, y int) int {
	return x + y
}

func sub(x float64, y float64) float64 {
	return x - y
}

func div(x float64, y float64) float64 {
	return x / y
}

func mult(y int) int {
	x := soma(5, 5)
	return x * y
}

func texto(nome string) string {
	return fmt.Sprintf("Olá, sou um programador iniciante em Go e me chamo %v.", nome)
}

// Se começar com letra maiuscula é uma função Publica
func Soma2(x, y int) (int, int) {
	return x + 1, y + 2
}
