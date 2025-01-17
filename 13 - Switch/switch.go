package main

import "fmt"

func diaDaSemana(num int) string {
	var diaDaSemana string

	switch num {
	case 1:
		diaDaSemana = "Domingo"
		fallthrough // vai jogar o código para a próxima condição
	case 2:
		diaDaSemana = "Segunda-Feira"
	case 3:
		diaDaSemana = "Terça-Feira"
	case 4:
		diaDaSemana = "Quarta-Feira"
	case 5:
		diaDaSemana = "Quinta-Feira"
	case 6:
		diaDaSemana = "Sexta-Feira"
	case 7:
		diaDaSemana = "Sábado"
	default:
		diaDaSemana = "Número Inválido."
	}

	return diaDaSemana
}

func diaDaSemana2(num int) string {
	switch {
	case num == 1:
		return "Domingo"
	case num == 2:
		return "Segunda-Feira"
	case num == 3:
		return "Terça-Feira"
	case num == 4:
		return "Quarta-Feira"
	case num == 5:
		return "Quinta-Feira"
	case num == 6:
		return "Sexta-Feira"
	case num == 7:
		return "Sábado"
	default:
		return "Número Inválido."
	}
}

func main() {
	fmt.Println("Switch")

	dia := diaDaSemana(1)
	fmt.Println(dia)

	dia2 := diaDaSemana2(1)
	fmt.Println(dia2)

}
