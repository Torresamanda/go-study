package main

import "fmt"

func main() {
	fmt.Println("Arrays Internos")

	slice := make([]float32, 10, 11)
	fmt.Println(slice)
	fmt.Println(len(slice)) //length
	fmt.Println(cap(slice)) //capacidade

	
}
