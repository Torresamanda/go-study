# Anotações

## Structs
 
| É uma coleção de campos, cada campo tem um nome e um tipo. 

Estrutura:

```bash
type usuario struct {
	nome string
	idade uint8
}
```

Existem três formas de utilizar a struct declarando os seus valores, primeira

```bash
var u usuario
u.nome = "Amanda"
u.idade = 25
fmt.Println(u)
```

Segunda

```bash
u2 := usuario{"Amanda", 25}
fmt.Println(u2)
```

Terceira, esta é utilizada para quando ainda não sabe todos os valores que serão preenchidos

```bash
u3 := usuario{nome: "Amanda"}
fmt.Println(u3)
```