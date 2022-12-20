package main

import "fmt"

func somar(n1 int16, n2 int16) int16 {
	return n1+n2
}

func calculos_matematicos(n1, n2 int16) (int16, int16){

	soma := n1 + n2
	subt := n1 - n2

	return soma, subt
}




func main(){
	soma := somar(10, 20)
	fmt.Println(soma)


	var f = func (txt string)  string {
		fmt.Println(txt)
		return txt
		
	}

	resultado := f("Passando o parametro da função f")
	fmt.Println(resultado)
	

	resultados_calc_soma, resultados_calc_sub := calculos_matematicos(10,15)
	fmt.Println(resultados_calc_soma," e ", resultados_calc_sub)

	resultados_calc_soma1, _ := calculos_matematicos(10,15)
	fmt.Println(resultados_calc_soma1)

	_ , resultados_calc_sub1 := calculos_matematicos(10,15)
	fmt.Println(resultados_calc_sub1)

}