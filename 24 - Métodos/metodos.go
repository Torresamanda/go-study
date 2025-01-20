package main

import "fmt"

type user struct {
	nome  string
	idade uint8
}

func (u user) salvar() {
	fmt.Printf("Salvando os dados do Usuário %s no banco de dados\n", u.nome)
}

func (u user) maiorDeIdade() bool {
	return u.idade >= 18
}

func (u *user) fazerAniversário() {
	u.idade++
}

func main() {
	user1 := user{"Amanda", 20}
	user1.salvar()

	user2 := user{"Davi", 45}
	user2.salvar()

	maiorDeIdade := user2.maiorDeIdade()
	fmt.Println(maiorDeIdade)

	user2.fazerAniversário()
	fmt.Println(user2.idade)
}
