package main

import "fmt"

func main() {
	var array [2]string

	array[0] = "Gustavo"
	array[1] = "Martins"

	numPrimos := [6]int{2, 3, 5, 7, 11, 13}

	fmt.Printf("%s %s", array[0], array[1])
	fmt.Println()
	fmt.Println(numPrimos)

	soma := numPrimos[0] + numPrimos[1]

	fmt.Println(soma)
	fmt.Println(numPrimos[0:4])

	slice := make([]string, 5)

	slice[0] = "Guzin"
	slice[1] = "Zika"

	fmt.Println(slice[0], slice[1])

	fmt.Println(len(slice))

	numPares := []int{2, 4, 6, 8}
	fmt.Println(numPares)
	fmt.Println(len(numPares))

	numPares = append(numPares, 10, 12)
	fmt.Println(numPares)
	fmt.Println(len(numPares))
}
