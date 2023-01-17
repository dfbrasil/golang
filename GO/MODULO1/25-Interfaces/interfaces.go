package main

import (
	"fmt"
	"math"
)


type retangulo struct {
	altura float64
	largura float64
}

func (r retangulo) area() float64 {
	return r.altura * r.largura
}

type circulo struct{
	raio float64
}

func (c circulo) area() float64{
	return math.Pi* math.Pow(c.raio,2)
}

//Declaração da INTERFACE
type forma interface{
	area() float64
}

func escrefver_area(f forma){

	fmt.Printf("A área da forma é %0.2f\n", f.area())

}


func main() {

	fmt.Println("Interfaces")

	ret := retangulo{10,15}
	escrefver_area(ret)
	
	cir := circulo{10}
	escrefver_area((cir))

}

// interfaces são utilizadas quando se precisa ter flexibilidade com os tipos