# Anotações

## Operadores

### Aritiméticos

- `+` (soma)
- `-` (subtração)
- `/` (divisão)
- `*` (multiplição)
- `%` (restos da divisão)

| No Go não é possível realizar operações aritiméticas com tipos diferentes, como no exemplo abaixo:

```bash
var num1 int16 = 10	
var num2 int32 = 25
somar := num1 + num2
```

### Atribuição

```bash
var variavel string = "String"
variavel2 := "String2"
fmt.Println(variavel, variavel2)
```

### Relacionais

| Sempre retornam um booleano (true or false)

- `>` (maior que)
- `>=` (maior, igual que)
- `==` (igual)
- `<=` (menor, igual que)
- `<` (menor que)
- `!=` (diferente)

### Lógicos

- `&&` (AND)
- `||` (OU)
- `!` (negação)

### Unários

```bash
numero := 10

numero++
numero--
numero += 15
numero -= 20
numero *= 3
numero /= 10
numero %= 2
	
fmt.Println(numero)
```

### Ternários

| Os operadores ternários não existem no GO por que a premissa do GO é que exista apenas uma maneira de fazer cada coisa, com exceção da declaração de variáveis. No GO se for necessário utilizar algo como o ternário utilizar if/else.
