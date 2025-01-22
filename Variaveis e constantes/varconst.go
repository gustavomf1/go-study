package main

import "fmt"

func main() {
	// Variáveis
	var nome string = "Gustavo"
	//nome = "Gustavo"

	fmt.Printf("Meu nome é %v Martins França\n", nome)

	numero := 2
	fmt.Println(numero)

	if numero == 1 {
		fmt.Printf("Meu nome é %v Martins França\n", nome)
	} else {
		fmt.Println("Sou louco!")
	}
}
