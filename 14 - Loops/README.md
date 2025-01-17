# Anotações

## Loops

### For

Estrutura:

```bash
i := 0

for i < 10 {
	i++
	fmt.Println("Incrementando i")
	time.Sleep(time.Second)
}

fmt.Println(i)
```

OU 

```bash
for j := 0; j < 10; j++ {
	fmt.Println("Incrementando j", j)
	time.Sleep(time.Second)
}
```

| A diferença entre os dois exemplos acima é que no primeiro exemplo a variável i está disponível para ser usando em qualquer parte da função, enquanto no segundo exemplo a variável j só está disponível dentro da condição.

### For com Range

Estrutura:

```bash
nomes := [3]string{"Amanda", "Tereza", "Leonardo"}

for indice, nome := range nomes {
	fmt.Println(indice, nome)
}
```

| O primeiro variável (no caso do exemplo acima "indice") é sempre o indice do array/slice, enquanto a segunda variável ("nome") é sempre o valor. Caso não queira utilizar o indice pode-se utilizar o "_" em seu lugar, desta forma:

```bash
nomes := [3]string{"Amanda", "Tereza", "Leonardo"}

for _, nome := range nomes {
	fmt.Println(nome)
}
```

### Interando em Strings

```bash
for indice, letra := range "PALAVRA" {
	fmt.Println(indice, string(letra))
}
```

| Obs: Necessário utilizar "string(variável)" para trazer cada letra ou o GO irá mostrar valores da tabela ASC

### Interando no map

```bash
user := map[string]string{
	"nome":      "Amanda",
	"sobrenome": "Torres",
}

for chave, valor := range user {
	fmt.Println(chave, valor)
}
```