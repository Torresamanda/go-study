# Anotações

## Switch

- Estrutura
```bash
switch num {
case 1:
    return "Domingo"
}
```

OU 

```bash
switch {
case num == 1:
    return "Domingo"
}
```

- fallthrough: é utilizada dentro do case em um bloco do switch, a declaração dela transfere o controle para o próximo case, independentemente de a condição ser verdadeira ou falsa. É uma palavra-chave do GO pouco utilizada, mas tem as suas utilidades.

```bash
switch {
case num == 1:
    return "Domingo"
    fallthrough
case num == 2:
    return "Segunda-Feira"
}
```

Obs: Não é necessário utilizar o "break" como nas outras linguagens, o GO já sai automaticamente do switch quando a condição é verdadeira.