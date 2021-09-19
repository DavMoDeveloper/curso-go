package main

import (
	"fmt"
)

func main() {
	animals := make(map[string]string)
	animals["cat"] = "Gato"
	animals["dog"] = "Perro"

	fmt.Println(animals)

	fruits := map[string]string{
		"apple":"manzana",
		"banana":"Platano",
	}
	fmt.Println(fruits)

	delete(fruits, "banana")
	fmt.Println(fruits)

	fmt.Println(animals["cat"])

	valor, ok := animals["gorilla"]
	fmt.Println(valor,ok)
}