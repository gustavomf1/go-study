package main

import "fmt"

func main() {
	type Pessoa struct {
		Nome  string
		Sexo  string
		Idade int
	}

	p1 := Pessoa{"Gustavo", "B", 21}

	if p1.Idade >= 18 {
		fmt.Printf("%s é de maior\n", p1.Nome)
	}

	if p1.Sexo == "M" {
		fmt.Printf("O %s é Homem!\n", p1.Nome)
	} else if p1.Sexo == "F" {
		fmt.Printf("O %s é Mulher!\n", p1.Nome)
	} else {
		fmt.Printf("O %s é indefinido!\n", p1.Nome)
	}

	if (p1.Idade % 2) == 0 {
		fmt.Printf("A idade de %s é par\n", p1.Nome)
	} else {
		fmt.Printf("A idade de %s é impar\n", p1.Nome)
	}

}
