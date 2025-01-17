# Anotações

## Estruturas de controle

### if e else

```bash
num := 10

if num > 15 {
	fmt.Println("Maior que 15")
} else {
    fmt.Println("Menor ou igual a 15")
}
```

### if init

| Ao utilizar o if init é possível declar um variável direto no if e utilizar apenas dentro da condição.

```bash
if num2 := num; num2 > 0 {
	fmt.Println("Número é maior que zero")
} else if num < - 10 {
	fmt.Println("Número é menor que - 10")
} else {
	fmt.Println("Entre 0 e -10")
}
```