package main

import "fmt"

func main(){

	soma := 1+2
	sub := 1-2
	div := 1/2
	mult := 1*2
	resto := 10%2

	fmt.Println(soma,sub,mult,div,resto)

	var num1 int16 = 10
	var num2 int32 =32
	var somar int16 = num1 + int16(num2)

	fmt.Println(somar)

	//ATRIBUIÇÃO

	var var1 string = "Teste"
	var2 := "String2"

	fmt.Println(var1, var2)

	// RELACIONAIS

	fmt.Println( 1 > 2)
	fmt.Println( 1 < 2)
	fmt.Println( 1 == 2)
	fmt.Println( 1 != 2)
	fmt.Println( 1 >= 2)
	fmt.Println( 1 <= 2)

	//LÓGICOS

	verdadeiro, falso := true ,false	
	fmt.Println(verdadeiro && falso)

	fmt.Println(verdadeiro || falso)

	fmt.Println(!verdadeiro)
	fmt.Print(!falso)

	//UNÁRIOS

	numero := 10

	numero++

	fmt.Println(numero)

	numero+=15

	fmt.Println(numero)

	numero--

	fmt.Println(numero)

	numero-=11

	fmt.Println(numero)
	
	numero *=2

	fmt.Println(numero)

	numero /=3

	fmt.Println(numero)

	numero%=2

	fmt.Println(numero)

	

}