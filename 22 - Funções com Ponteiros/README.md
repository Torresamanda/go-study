# Anotações

## Funções com Ponteiros

| Um ponteiro (ou apontador) nada mais é do que uma variável que, ao invés de armazenar um valor (true, "hello world"), ela armazena um endereço que está alocado na memória. 

### Estrutura

```bash
func inverterSinalPonteiro(numero *int) {
	*numero = *numero * -1
}

func main() {
	novoNumero := 40
	fmt.Println(novoNumero)
	inverterSinalPonteiro(&novoNumero)
	fmt.Println(novoNumero)
}
```