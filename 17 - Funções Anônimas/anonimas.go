package main

import "fmt"

func main() {
	//Simples
	func() {
		fmt.Println("Olá mundo")
	}()

	//Utilizando parâmetrocs
	func(text string) {
		fmt.Println(text)
	}("Passando o parâmetro")

	//Utilizando o retorno
	retorno := func(text string) string {
		return fmt.Sprintf("Recebido -> %s", text)
	}("Passando o parâmetro")

	fmt.Println(retorno)
}
