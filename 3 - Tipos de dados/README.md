# Anotações

## Tipos de dados

#### Números

- Int: A diferença entre eles é a quantidade de bites que o número pode ter, ou seja seu tamanho. O Int sozinho vai utilizar a arquitetura do computador como base, por exemplo se o computador for 64 bites ele utilizará o `int64`.
    - `int8`
    - `int16`
    - `int32` ou `rune`
    - `int64` 

- Uint: Unsigned, que significa sem sinal, não assinalado. É um Int que desconsidera o sinal. 
    - `uint8` ou `byte`
    - `uint16`
    - `uint32`
    - `uint64` 

```bash
var numero2 uint32 = -2000 // Este é um exemplo de como não utilizar o uint
```
- Float: Segue a mesma lógica dos int, com a difença que ao utilizar o float ele irá suportar o ".", ou seja, número quebrados.
    - `float32`
    - `float64`

### Strings

Obs: 

 - O go não tem os caracteres Char, ele é convertido para um número. Os chars são declarados com aspas simples ('').

- Sempre aspas duplas ("") para declarar variáveis do tipo string 

### Valor zero

| No Go todo tipo de dado tem um valor, no caso dos números é o valor zero (0) para string é um valor vazio (uma string vazia) e para erro é o nil. 

### Boolean

Obs: O Valor zero do boolean é "false"

### Error

| O Go tem um pacote interno para realizar as tratativas de erros

Obs: O Valor zero do boolean é "nil"
