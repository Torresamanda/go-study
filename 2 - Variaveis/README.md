# Anotações

## Variáveis

| O Go é uma linguagem fortemente tipagada, dito isso é importante conhecer os tipos e como declara as variáveis. Existem duas maneiras de declara uma variável, a primeira delas é deixando o tipo de ela explícito e a segunda é deixando o tipo implícito.


### ❗Importante!

- O Go não permite importar e criar variáveis que não serão utilizados, ele não compilará o código.

```bash
	var variavel1 string = "Variável 1" //explícita
	variavel2 := "Variável 2" //implícita -. inferência de tipo
```

| O Go permite declarar várias variáveis de maneira exlícita, conforme o exemplo abaixo:

```bash
var (
	variavel3 string = "etc etc etc"
	variavel4 string = "etc etc etc"
)
```

| E, também, de maneira implícita:


```bash
variavel5, variavel6 := "Variavel 5", "Variavel 6"
```

Obs: Nos exemplos acima foi utilizado variáveis, mas o mesmo conceito se aplica as constantes. 

### Curiosidade do Go

- É possível, no Go, inverter os valores das variáveis, desta forma:

```bash
	variavel5, variavel6 = variavel6, variavel5
``` 