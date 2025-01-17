# Anotações

## Map

| Dentro dos colchetes é o tipo das chaves e fora do colchetes é o tipo dos valores.

```bash
map[string]string 
```

- Como utilizar:

```bash
user := map[string]string{
	"nome":      "Pedro",
	"sobrenome": "Silva",
}

fmt.Println(user["nome"])
```

- Para deletar um valor da chave:

```bash
delete(user2, "nome")
fmt.Println(user2)
```

- Para adicionar uma chave e valor novo:

```bash
user2["signo"] = map[string]string{
	"signo": "Gêmeros",
}
	
fmt.Println(user2)
```