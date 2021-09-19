package main

import (
	"fmt"
)

func main(){
	//Operadores Aritméticos (), *, /, %, +; -
	var a = 4 + 2*3
	fmt.Println(a)

	//Operadores de asignación: =, +=, -=, *=, /=, %=
	var b = 10
	b += 2
	fmt.Println(b)

	//declaración post-incremento y post-decremento: ++, --
	//(no son una expresión sino una declaración)
	var c = 3
	c++
	fmt.Println(c)
	c--
	fmt.Println(c)

	//Operadores Compración: >, <, ==, != >=, <=
	fmt.Println(4>5)
	fmt.Println(4<5)
	fmt.Println(4==4)
	fmt.Println(4!=4)

	//Operadores Lógicos &&, || 
	var age = 30
	fmt.Println("Adulto:" , age >= 18 && age <= 60)
	fmt.Println("Niño o viejo:", age < 18 && age > 60)
	//Unario: !
	fmt.Println("Operador Uranio niega el valor",!(4 == 4))

}