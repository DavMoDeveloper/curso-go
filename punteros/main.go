package main

import (
	"fmt"
)

func main() {
	fruit := "Manzana"
	var puntero *string 
	puntero = &fruit //direccion en memoria
	fmt.Printf("Tipo: %T, Valor: %s, Dirección %v\n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", puntero, puntero, *puntero)

	*puntero = "Piña"
	fmt.Printf("Tipo: %T, Valor: %s, Dirección %v\n", fruit, fruit, &fruit)
	fmt.Printf("Tipo: %T, Valor: %v, Desreferenciación: %s\n", puntero, puntero, *puntero)
}