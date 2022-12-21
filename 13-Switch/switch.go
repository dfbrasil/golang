package main

import (
	"fmt"
)

func dia_da_semana(numero int) string{
	switch numero{
	case 1:
		return "Domingo"
	case 2:
		return "Segunda"
	case 3:
		return "Terça"
	case 4:
		return "Quarta"
	case 5:
		return "Quinta"
	case 6:
		return "Sexta"
	case 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}

}

func dia_da_semana2(numero int) string{

	switch{
	case numero == 1:
		return "Domingo"
	case numero == 2:
		return "Segunda"
	case numero == 3:
		return "Terça"
	case numero == 4:
		return "Quarta"
	case numero == 5:
		return "Quinta"
	case numero == 6:
		return "Sexta"
	case numero == 7:
		return "Sábado"
	default:
		return "Número Inválido"
	}

}

func dia_da_semana3(numero int) string{

	var dia_da_semana string

	switch{
	case numero == 1:
		dia_da_semana = "Domingo"
	case numero == 2:
		dia_da_semana = "Segunda"
	case numero == 3:
		dia_da_semana = "Terça"
	case numero == 4:
		dia_da_semana = "Quarta"
	case numero == 5:
		dia_da_semana = "Quinta"
	case numero == 6:
		dia_da_semana = "Sexta"
	case numero == 7:
		dia_da_semana = "Sábado"
	default:
		dia_da_semana = "Número Inválido"
	}
	return dia_da_semana
}

func dia_da_semana4(numero int) string{

	var dia_da_semana string

	switch{
	case numero == 1:
		dia_da_semana = "Domingo"
		fallthrough
	case numero == 2:
		dia_da_semana = "Segunda"
	case numero == 3:
		dia_da_semana = "Terça"
	case numero == 4:
		dia_da_semana = "Quarta"
	case numero == 5:
		dia_da_semana = "Quinta"
	case numero == 6:
		dia_da_semana = "Sexta"
	case numero == 7:
		dia_da_semana = "Sábado"
	default:
		dia_da_semana = "Número Inválido"
	}
	return dia_da_semana
}

func main(){

	fmt.Println("Switch")

	dia := dia_da_semana(51)
	fmt.Println(dia)

	dia2 := dia_da_semana2(3)
	fmt.Println(dia2)

	dia3 := dia_da_semana3(5)
	fmt.Println(dia3)

	dia4 := dia_da_semana4(1) // FALLTHOUGTH
	fmt.Println(dia4)

}