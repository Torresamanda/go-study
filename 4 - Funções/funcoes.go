package main

import "fmt"

func somar(n1 int8, n2 int8) int8 {
	return n1 + n2
}

func calculosMatematicos(n1, n2 int8) (int8, int8) {
	soma := n1 + n2
	sub := n1 - n2
	return soma, sub
}

func main() {
	somar := somar(10, 20)
	fmt.Println(somar)

	var f = func(txt string) {
		fmt.Println(txt)
	}

	f("Hello Word")

	resultadosCalculos, _ := calculosMatematicos(10, 15)
	fmt.Println(resultadosCalculos)
}