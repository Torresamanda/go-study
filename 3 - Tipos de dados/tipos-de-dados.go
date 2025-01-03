package main

import "fmt"

func main() {
	var numero int64 = 1000000000000
	fmt.Println(numero)

	var numero2 uint32 = 2000
	fmt.Println(numero2)

	//alias
	var numero3 rune = 12345
	fmt.Println(numero3)

	var numero4 byte = 123
	fmt.Println(numero4)

	var numeroReal1 float32 = 123.54
	fmt.Println(numeroReal1)

	var numeroReal2 float64 = 13423535223.54
	fmt.Println(numeroReal2)

	// Fim dos números

	var str string = "Texto"
	fmt.Println(str)

	char := 'B'
	fmt.Println(char)

	// fim das strings

	var boolean1 bool
	fmt.Println(boolean1)

	var erro error
	fmt.Println((erro))
}
