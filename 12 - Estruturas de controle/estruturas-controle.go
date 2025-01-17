package main

import "fmt"

func main() {
	fmt.Println("Estruturas de controle")

	num := 10

	if num > 15 {
		fmt.Println("Maior que 15")
	} else {
		fmt.Println("Menor ou igual a 15")
	}

	if num2 := num; num2 > 0 {
		fmt.Println("Número é maior que zero")
	} else if num < - 10 {
		fmt.Println("Número é menor que - 10")
	} else {
		fmt.Println("Entre 0 e -10")
	}
}
