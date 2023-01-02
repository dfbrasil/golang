package main

import "fmt"

func main(){

	fmt.Println("Ponteiros")

	var var1 int = 10
	var var2 int = var1
	fmt.Println(var1, var2)

	var1++
	fmt.Println(var1, var2)

	//Ponteiro é uma referencia de memória

	var var3 int
	var ptr *int
	fmt.Println(var3, ptr)

	var3 =100
	ptr = &var3
	fmt.Println(var3, ptr)

	//desreferenciação
	fmt.Println(*ptr)

	var3 = 150
	fmt.Println(var3, ptr)
	fmt.Println(var3, *ptr)


}