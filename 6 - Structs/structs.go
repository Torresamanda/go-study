package main

import "fmt"

type usuario struct {
	nome  string
	idade uint8
}

func main() {
	fmt.Println("Arquivos Structs")

	var u usuario
	u.nome = "Amanda"
	u.idade = 25
	fmt.Println(u)

	u2 := usuario{"Amanda", 25}
	fmt.Println(u2)

	u3 := usuario{nome: "Amanda"}
	fmt.Println(u3)
}
