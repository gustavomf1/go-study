package main

import "fmt"

type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    float32
}

type estudante struct {
	pessoa
	curso     string
	faculdade string
}

func main() {
	p1 := pessoa{"Gustavo", "Magnossão", 22, 1.65}
	fmt.Println(p1)

	e1 := estudante{p1, "Sistema de informação", "Toledo"}

	fmt.Println(e1)
	fmt.Println(e1.idade)
}
