package main

import "fmt"

func main() {
	var variavel1 string = "Variável 1" //explicita
	variavel2 := "Variável 2"           //implicita -. inferencia de tipo

	fmt.Println(variavel1)
	fmt.Println(variavel2)

	var (
		variavel3 string = "etc etc etc"
		variavel4 string = "etc etc etc"
	)

	fmt.Println(variavel3, variavel4)

	variavel5, variavel6 := "Variavel 5", "Variavel 6"

	fmt.Println(variavel5, variavel6)

	variavel5, variavel6 = variavel6, variavel5
	fmt.Println(variavel5, variavel6)
}
