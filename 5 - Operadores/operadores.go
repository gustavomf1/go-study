package main

import "fmt"

func main() {
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 1 / 2
	multiplicacao := 1 * 2
	restoDaDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDaDivisao)

	var variavel string = "String"

	variavel2 := "String2"
	fmt.Println(variavel, variavel2)

	fmt.Println(2 > 1)
	fmt.Println(2 >= 1)
	fmt.Println(2 == 1)
	fmt.Println(2 <= 1)
	fmt.Println(2 != 1)
	fmt.Println(2 < 1)

	if soma == 0 && subtracao == 0 {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}

	if soma == 3 || subtracao == 0 {
		fmt.Println(!true)
	} else {
		fmt.Println(false)
	}
}
