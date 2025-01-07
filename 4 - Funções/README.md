# Anotações

## Funções

- Retorno da função é depois do fechamento dos parênteses `func somar(a int, b int) int`

Exemplo de uso retornando duas funções:

```bash
func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}
```

- No Go é possível retornar dois valores na mesma função, como no exemplo acima, caso queira ignora o primeiro/segundo valor de retorno é feito conforme o exemplo abaixo, considere que os cálculos ainda é feito, mas o resultado é ignorado. 

```bash
resultadosCalculos, _ := calculosMatematicos(10, 15)
fmt.Println(resultadosCalculos)
```
