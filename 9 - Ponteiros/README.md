# Anotações

## Ponteiros

| Pode-se pensar no ponteiro como uma variável que irá salvar, não um valor, mas um endereço de memória de alguma coisa.

```bash
// Ponteiro é uma referência de memória
var var3 int
var ponteiro *int // A variável ponteiro guarda um endereço de memória de um valor inteiro

var3 = 100
ponteiro = &var3 

fmt.Println(var3, *ponteiro) //desreferenciação
```