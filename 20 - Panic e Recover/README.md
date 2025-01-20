# Anotações

## Panic 

| É uma função usada para disparar um pânico de tempo de execução. Quando o panic ocorre, o fluxo normal do programa é interrompido, e o programa começa a desenrolar a pilha de chamadas, executando quaisquer funções adiadas ao longo do caminho. Eles servem para sinalizar erros inrrecuperáveis ou condições excepcionais.

## Recover

| Para lidar com o panic e se recuperar deles, o Golang fornece a função recover(). Quando usada em conjunto com defer permite capturar o valor do panic e retomar a função execução normal. Se estivermos em modo de panic, a chamada para recover() retorna o valor passado para a função panic. 

### Estrutura:

```bash
func recuperarExecucao() {
	if r := recover(); r != nil {
		fmt.Println("Execução recuperada com sucesso!")
	}
}

func alunoEstaAprovado(n1, n2 float64) bool {
	defer recuperarExecucao()
	media := (n1 + n2) / 2

	if media > 6 {
		return true
	} else if media < 6 {
		return false
	}

	panic("A média é exatamente 6!")
}

func main() {
	fmt.Println(alunoEstaAprovado(6, 6))
	fmt.Println("Pós execução!")

}
```