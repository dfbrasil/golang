package main

import (
	"errors"
	"fmt"
)

func main(){

	var num int16 = 1000
	fmt.Println(num)

	var num2 uint16 = 2 //unsygned int
	fmt.Println(num2)

	// allias
	var num3 rune = 123 // rune == int32
	fmt.Println(num3)

	var num4 byte = 8 // byte ==int8
	fmt.Println(num4)
	
	var real1 float32 = 12.6
	fmt.Println(real1)

	var real2 float64 = 123.66
	fmt.Println(real2)

	real3 := 1324.56
	fmt.Println(real3)

	var str string = "texto de str"
	fmt.Println(str)

	str2 := "texte de inferecia"
	fmt.Println(str2)

	char := 'D'
	fmt.Println(char)

	//Valor 0

	var valor0num int
	fmt.Println(valor0num)

	var valor0str string
	fmt.Println(valor0str)

	var bool0bool bool
	fmt.Println(bool0bool)

	var bool1 bool = true
	fmt.Println(bool1)

	var erro0error error
	fmt.Println(erro0error)

	var erro error = errors.New("Erro interno")
	fmt.Println(erro)

}