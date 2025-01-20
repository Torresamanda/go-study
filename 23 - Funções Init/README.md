# Anotações

## Funções Init

| Função que será executada antes da função `main`. É possível ter uma função `init` por arquivo, usada para fazer setup, variaveis que serão iniciadas, etc.

### Estrutura

```bash
var n int

func init() {
	fmt.Println("Executando a função init")
	n = 10
}

func main() {
	fmt.Println("Executando a função main")
	fmt.Println(n)
}
```