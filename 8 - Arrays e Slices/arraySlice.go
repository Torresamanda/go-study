package main

import (
	"fmt"
	"reflect"
)

func main() {
	fmt.Println("Arrays e Slices")

	var array1 [5]int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [2]string{"Posição1", "Posição2"}
	fmt.Println(array2)

	array3 := [...]int{1, 2, 3, 4, 5}
	fmt.Println(array3)

	slice1 := []int{10, 11, 12, 13, 14, 15, 16}
	fmt.Println(slice1)

	fmt.Println((reflect.TypeOf(slice1)))
	fmt.Println((reflect.TypeOf(array3)))

	slice1 = append(slice1, 44)
	fmt.Println(slice1)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[1] = 21
	fmt.Println(slice2)
}
