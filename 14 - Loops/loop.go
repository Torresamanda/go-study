package main

import (
	"fmt"
)

func main() {
	fmt.Println("Loops")

	// i := 0

	// for i < 10 {
	// 	i++
	// 	fmt.Println("Incrementando i")
	// 	time.Sleep(time.Second)
	// }

	// fmt.Println(i)

	// for j := 0; j < 10; j += 2 {
	// 	fmt.Println("Incrementando j", j)
	// 	time.Sleep(time.Second)
	// }

	// nomes := [3]string{"Amanda", "Tereza", "Leonardo"}

	// // o primeiro por padrão é sempre o indice e o segundo é sempre o valor.
	// for _, nome := range nomes {
	// 	fmt.Println(nome)
	// }

	// for indice, letra := range "PALAVRA" {
	// 	fmt.Println(indice, string(letra))
	// }

	user := map[string]string{
		"nome":      "Amanda",
		"sobrenome": "Torres",
	}

	for chave, valor := range user {
		fmt.Println(chave, valor)
	}
}
