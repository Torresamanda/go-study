# Anotações

## Funções Closure 

| Clousure são recursos que permite encapsular estado e comportamento dentro de funções. Eles fornecem uma maneira de criar unidades de códigos autocontidas que podem acessar e manipular variáveis de seu escopo circundante. 

### Estrutura

```bash
func closure() func() {
	texto := "Dentro da função closure"

	funcao := func() {
		fmt.Println(texto)
	}

	return funcao
}

func main() {
	texto := "Dentro da função main"
	fmt.Println(texto)

	funcaoNova := closure()
	funcaoNova()
}
```