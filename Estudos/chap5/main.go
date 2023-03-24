package main

import (
	"fmt"
	"time"
)

func main(){
	var hello string = "olá mundo!"
	fmt.Printf("thats my hello! %v \n", hello)
	fmt.Println("hora e data abaixo")
	diaHora := time.Now()
	fmt.Printf("Dia e hora: %v - %v \n", diaHora.Format("02/01/2006 15:04:05"), diaHora.Weekday())
}
