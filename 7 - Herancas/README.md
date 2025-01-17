# Anotações

## "Heranças"

| É o mais perto que o GO chega das heranças convencionais

```bash
type pessoa struct {
	nome      string
	sobrenome string
	idade     uint8
	altura    uint8
}

type estudante struct {
	pessoa //único caso que não precisa especificar o tipo.
	curso string
	faculdade string
}
```

Como utilizar:

```bash
p1 := pessoa{"João", "Pedro", 20, 178}
fmt.Println(p1)

e1 := estudante{p1, "engenharia", "USP"}
fmt.Println(e1)
```