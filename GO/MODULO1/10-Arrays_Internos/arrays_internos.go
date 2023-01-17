package main

import (
	"fmt"
	"reflect"
)

func main()  {
	fmt.Println("Arrays Internos")

	var array1[5] int
	array1[0] = 1
	fmt.Println(array1)

	array2 := [5] string{"posição1","posição2"}
	fmt.Println(array2)

	array3 := [...]int{1,2,3,4,5}
	fmt.Println(array3)

	slice1 := []int {1,2,3,4,5,6}
	fmt.Println(slice1)

	fmt.Println(reflect.TypeOf(slice1))
	fmt.Println(reflect.TypeOf(array3))

	slice1 = append(slice1, 22)
	fmt.Println(slice1)

	slice2 := array3[1:3]
	fmt.Println(slice2)

	array3[1] = 567
	fmt.Println(slice2)

	//ARRAYS INTERNOS

	slice3 := make([]float32, 10, 11)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice3 = append(slice3, 22)
	fmt.Println(slice3)

	slice3 = append(slice3, 33)
	fmt.Println(slice3)
	fmt.Println(len(slice3))
	fmt.Println(cap(slice3))

	slice4 :=make([]float32, 5) // se não passar o ultimo parâmetro (cap) o slice vai setar o cap igual ao len
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

	slice4 = append(slice4, 555)
	fmt.Println(slice4)
	fmt.Println(len(slice4))
	fmt.Println(cap(slice4))

}
