# Anotações

## Arrays e Slices

### Arrays

```bash
var array1 [5]int
fmt.Println(array1)
```

| O GO não permite criar um array de tipos diferentes, ou ele é int ou string, etc. É obrigatório especificar o tamanho do array.

- Como difinir as posições do array?


```bash
var array1 [5]int
array1[0] = 1
fmt.Println(array1)
```

OU

```bash
array2 := [2]string{"Posição1", "Posição2"}
fmt.Println(array2)
```

| Tentar inserir mais valores em um array com um tamanho definido não é possível tornando os array inflexíveis. Essa é uma maneira de deixa-lo um pouco mais flexível, mas ainda não é permitido adicionar valores posteriormente, sendo uma prática pouco utilizada

```bash
array3 := [...]int{1, 2, 3, 4, 5}
fmt.Println(array3)
```

### Slices

| O slice é uma estrutura baseada no array, mas o slice é flexível principalmente na questão de tamanho. O Slice também tem a limitação de não poder ter tipos diferentes dentro dele.

```bash
slice1 := []int{10, 11, 12, 13, 14, 15, 16}
fmt.Println(slice1)
```

| Importante ressaltar que o slice não é um array apesar de ser parecido, o slice aponta para o array, como se fosse uma fatia de um array.

#### Append

| A função append vai pegar o slice que foi referenciado e adicionar o novo valor especificado, no caso do exemplo abaixo a função está pegando os valores já adicionado no slice e adicionando um novo valor.

```bash
slice1 = append(slice1, 44)
fmt.Println(slice1)
```

### Diferença entre Arrays e Slices

```bash
fmt.Println((reflect.TypeOf(slice1)))
fmt.Println((reflect.TypeOf(array3)))

//Resultado
//[]int
//[5]int
```

### Referenciar um Array

```bash
slice2 := array3[1:3] // posição do array 1 e 2, o resultado será os valores que estão nesta posição dentro do array
fmt.Println(slice2)
```