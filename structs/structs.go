package main

import (
	"fmt"
)

type Pessoa struct {
	Nome  string
	Idade int
}

func main() {
	fmt.Println(Pessoa{"Steph", 22})

	p1 := Pessoa{"Gustavo", 21}

	fmt.Println(p1)

	p2 := Pessoa{Nome: "Zezin", Idade: 12}

	fmt.Println(p2)

	fmt.Println(p1.Nome, p2.Idade)

	pessoas := []Pessoa{}
	pessoas = append(pessoas, p1, p2)

	fmt.Println(pessoas)

	classe := map[string][]Pessoa{}

	classe["Sistemas de informação"] = pessoas
	classe["Teste"] = pessoas
	fmt.Println(classe)
}
