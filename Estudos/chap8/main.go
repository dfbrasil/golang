package main

import (
	
	"fmt"
	"time"
)

const hotelName string = "Hotel California"
const longitude = 24.806078
const latitude = -78.243027
var ocuppancy int = 12

func Hotel() {
	fmt.Printf("O nome do Hotel é: %s\nLongitude: %f\nLatitude: %f\nOcupação: %d\n", hotelName, longitude, latitude, ocuppancy)
}

// slice of bytes
func main(){
	b := make([]byte, 0)
	b = append(b, 255)
	b = append(b, 10)
	b = append(b, 22)
	fmt.Printf("o slice de bites é: %v \n",b)

	palavra := "teste"

	for i, c := range palavra{
		fmt.Printf("a letra %d é %s\n", i, string(c))
	}

	loc, err := time.LoadLocation("America/Sao_Paulo")
	if err != nil{
		fmt.Println(err)
	}
	fmt.Println(loc)

	Hotel()

}
