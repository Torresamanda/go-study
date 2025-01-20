# Anotações

## Métodos

| Os métodos são obrigatoriamente associado alguma coisa, uma struct, interface, etc.

### Estruturas

```bash
type user struct {
	nome  string
	idade uint8
}

func (u user) salvar() {
	fmt.Println("Dentro do método salvar")
}

func main() {
	user1 := user{"Usuário 1", 20}
	fmt.Println(user1)
	user1.salvar()
}

```