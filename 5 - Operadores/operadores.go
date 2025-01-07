package main

import "fmt"

func main() {
	// Aritiméticos
	soma := 1 + 2
	subtracao := 1 - 2
	divisao := 10 / 5
	multiplicacao := 10 * 5
	restoDivisao := 10 % 2

	fmt.Println(soma, subtracao, divisao, multiplicacao, restoDivisao)

	var num1 int16 = 10
	var num2 int16 = 25
	somar := num1 + num2
	fmt.Println(somar)

	// Atribuição
	var variavel string = "String"
	variavel2 := "String2"
	fmt.Println(variavel, variavel2)

	// Relacionais
	fmt.Println(1 > 2)
	fmt.Println(1 >= 2)
	fmt.Println(1 == 2)
	fmt.Println(1 <= 2)
	fmt.Println(1 < 2)
	fmt.Println(1 != 2)

	// Lógicos
	fmt.Println(true && false)
	fmt.Println(true || false)
	fmt.Println(!true)

	// Unários
	numero := 10

	numero++
	numero--
	numero += 15
	numero -= 20
	numero *= 3
	numero /= 10
	numero %= 2

	fmt.Println(numero)
}
