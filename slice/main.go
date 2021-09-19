package main

import (
	"fmt"
)

func main() {
	//Slivce no poseen datos, son apuntadores a un Array.
	set := [7]string{"Leon","Caballo","Perro","Mariposa","Loro","Avion","Martillo"}
	animals := set[0:5]
	fly := set[3:7] 
	fly[0] = "Aguila"
	fmt.Println("Array:", set)
	fmt.Println("Animales:", animals)
	fmt.Println("Voladores:", fly)
	fmt.Println("All:", set[:])
	//len(): # de elementos en el slice
	//cap(): # de elementos del array donde apunta el slice, a
	//partir del indice de donde se creo el slice
	fmt.Println("Tamaño:", len(animals))
	fmt.Println("Capacida:", cap(animals))
	//Agregar elemntos a un slice
	animals = append(animals,"Gato")
	fmt.Println("Array:", set)
	fmt.Println("Animales:", animals)

	//Creacion de slice a partir de un array
	//fruits := []string{"Fresa","Melon"}
	//Creación de un array atraves de un slice
	fruits := make([]string, 0, 3)
	fruits = append(fruits, "Piña")
	fmt.Println("fruits:", fruits)
	fmt.Println("tamaño:",len(fruits))
	fmt.Println("capacidad", cap(fruits))
}