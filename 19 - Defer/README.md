# Anotações

## Defer

| É uma palavra reservada do GO que envia a função para uma lista de chamadas. Essas chamadas serão executadas depois que a função onde o defer se encontra terminar a sua execução. O defer é muito utilizado para realizar operações de limpeza, como por exemplo, fechar conexões com banco de dados, fechar stream de arquivos e etc.

### Estrutura

```bash
func alunosEstaAprovados(n1, n2 float32) bool {
	defer fmt.Println("Média calculada. Resultado será retornado.")
	fmt.Println("Entrando na função para verificar se o aluno está aprovado.")

	media := (n1 + n2) / 2

	if media >= 6 {
		return true
	}

	return false
}

func main() {
	fmt.Println(alunosEstaAprovados(2, 5))
}
```