package main

import (
	"fmt"
)

func Soma(valores ...int) (total int){
	for _, valor := range valores{
		total += valor
	}
	return
}

func JustPanic(b bool) {
    if b {
        panic("Fire!!! Fire!!! Fire!!!")
    }
}

func main()  {
	fmt.Println("Pasta de testes")
}