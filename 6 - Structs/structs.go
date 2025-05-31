package main

import "fmt"

type usuario struct {
	nome     string
	idade    int8
	endereco endereco
}

type endereco struct {
	logradouro string
	numero     int8
}

func main() {
	fmt.Println()

	var u usuario
	u.nome = "Gustavo"
	u.idade = 22
	fmt.Println(u)

	enderecoExemplo := endereco{"Rua zé", 2}
	u2 := usuario{"Zé", 23, enderecoExemplo}
	fmt.Println(u2)

	u3 := usuario{nome: "Teste"}
	fmt.Println(u3)
}
