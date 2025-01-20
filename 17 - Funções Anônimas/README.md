# Anotações

## Funções Anônimas

| Uma função anônima é uma função que não tem nome. Ela é útil quando você quer criar uma função inline. Em GO, uma função anônima também pode formar um fechamento. Uma função anônima também é conhecida como literal de função.

### Estrutura:

```bash
// Funçãoa anônima simples
func main() {
	func() {
		fmt.Println("Olá, Mundo!")
	}()
}
```