package main

import "fmt"

func main() {
	notas := map[string]float64{}
	notas["Meg"] = 7
	notas["Ze"] = 10
	notas["Gordin"] = 8
	notas["Mar"] = 1

	media := (notas["Meg"] + notas["Ze"] + notas["Gordin"] + notas["Mar"]) / 4

	fmt.Println(media)

	anoNas := map[int]int{
		1: 2003,
		2: 1000,
	}

	fmt.Println(anoNas[1])

	anoNas[3] = 2001
	fmt.Println(anoNas)

	trueOrFalse := map[bool]bool{
		true:  false,
		false: true,
	}

	fmt.Println(trueOrFalse[false])
}
