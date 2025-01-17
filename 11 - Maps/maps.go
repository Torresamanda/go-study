package main

import "fmt"

func main() {
	fmt.Println("Maps")

	user := map[string]string{
		"nome":      "Pedro",
		"sobrenome": "Silva",
	}

	fmt.Println(user["nome"])

	user2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Amanda",
			"segundo":  "Torres",
		},
	}

	fmt.Println(user2)
	fmt.Println(user2["nome"]["primeiro"])

	delete(user2, "nome")
	fmt.Println(user2)

	user2["signo"] = map[string]string{
		"signo": "Gêmeros",
	}
	
	fmt.Println(user2)
}
