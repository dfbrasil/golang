package main

import (
	"fmt"
	"math/rand"
	"time"
)

// var hotelName string = "Hotel California"
// var roomsAvailable, rooms, roomsOccupied uint8


func idade() {
	rand.Seed(time.Now().UTC().UnixNano())
	var ageJonh, agePaul uint8 =  uint8(rand.Intn(100)), uint8(rand.Intn(100))
	fmt.Printf("Jonh tem %v anos.\n", ageJonh)
	fmt.Printf("Paul tem %v anos.\n", agePaul)
	if ageJonh > agePaul{
		fmt.Println("É verdade que João é mais velho que Paul")
	} else if ageJonh < agePaul{
		fmt.Println("É verdade que Paulo seja mais velho que Jonh.")
	} else{
		fmt.Println("Ambos tem a mesma idade.")
	}
}

func main()  {
	// fmt.Printf("O nome do hotel é: %v\n", hotelName)
	// fmt.Println("Insira o número de quartos:")
	// fmt.Scanf("%v\n", &rooms)
	// fmt.Println("Insira o número de quartos ocupados:")
	// fmt.Scanf("%v\n", &roomsOccupied)
	// if roomsOccupied > rooms{
	// 	fmt.Println("Número de quartos ocupados maior que o disponível")
	// }
	// roomsAvailable = rooms - roomsOccupied
	// if roomsAvailable == 0{
	// 	fmt.Println("Não há leitos disponíveis")
	// } else{
	// 	fmt.Printf("Existem %v disponíveis \n", roomsAvailable )
	// }
	idade()
}