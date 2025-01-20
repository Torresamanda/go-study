# Anotações

## Funções Recursivas

| É uma função que chama a si mesma. As funções recursivas podem ser muito úteos para resolver problemas que podem ser divididos em subproblemas menores e semelhanetes, como calcular o fatorial de um número, gerar a sequência de Fibonacci ou percorrer uma estrtura de dados em árvore. 

Obs: A sequência de Fibonacci é uma sequência de números onde o próximo número é a soma dos dois anteriores `1 1 2 3 5 8 13`, ela pode começar com 0 ou 1.

### Estrutura:

```bash
func fibonacci(position uint) uint {
	if position <= 1 {
		return position
	}

	return fibonacci(position-2) + fibonacci(position-1)
}

func main() {

	position := uint(10)
	fmt.Println((fibonacci(position)))
}

```

| Nas funções recursivas é obrigatório ter o momento de parada da função, para não gerar um estouro de pilha (stackoverflow).

Obs: As funções recursivas não são recomendas em execuções muito grande.  