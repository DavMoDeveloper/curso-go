package main

import "fmt"

func main(){
	// bool, string, numeric
	var a bool = true
	
	fmt.Printf("Tipo: %T, Valor %v", a, a)

	fmt.Println()

	var b string = "EDTeam" 

	fmt.Printf("Tipo: %T, Valor %v", b, b)

	println()

	var edad uint8 = 200 

	fmt.Printf("Tipo: %T, Valor %v", edad, edad)

	println()

	//byte = uint8
	var numero byte = 255

	fmt.Printf("Tipo: %T, Valor %v", numero, numero)

	println()

	//rune = uint32 and rune = unicode
	var numero2 rune = -100 

	fmt.Printf("Tipo: %T, Valor %v", numero2, numero2)

	println()

	var unicode rune = 'a'
	fmt.Printf("Tipo: %T, Valor %v", unicode, unicode)

	println()

	//Flotantes
	var flotante float64 = 23.24
	fmt.Printf("Tipo: %T, Valor %v", flotante, flotante)

	//Operador Black definir una variable sin usarla
	var _ string  = "test"

}