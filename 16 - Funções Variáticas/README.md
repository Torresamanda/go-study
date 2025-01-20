# Anotações

## Funções Variáticas

| É uma função que pode receber any parâmetros, não é necessário especificar a os parâmetros que ela vai receber. Sendo possível não passar nenhum um parâmetro também, só lembrando não ser posível passar um parâmetro de um tipo diferente do especificado.

### Estrutura:

```bash
func soma(numeros ...int) int {
	total := 0
	for _, numero := range numeros {
		total += numero
	}

	return total
}
```