# Anotações

## Arrays Internos

| Um Array tem uma tamanho fixo, enquanto o slice não. O slice quando tem a sua capacidade excedida automaticamente o GO dobra esse valor, tornando ele um pouco mais "flexível".

```bash
slice := make([]float32, 10, 11)
fmt.Println(slice)
fmt.Println(len(slice)) //length
fmt.Println(cap(slice)) //capacidade
```

| A função make utilizada no exemplo acima, é usada para inicializar fatias, mapas e canais - estrutras de dados que exigem inicialização em tempo de execução.