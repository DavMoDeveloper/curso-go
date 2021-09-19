package main

import (
	"fmt"
)

func main() {

	variable := "gato"
	/*switch variable {
	case "gato", "perro":
		fmt.Println("Es un animal")
	case "platano", "fresa":
		fmt.Println("Es una fruta")		
	default:
		fmt.Println("no es un animal ni fruta")
	}*/

	switch  {
	case variable == "gato" || variable == "perro":
		fmt.Println("Es un animal")
	case variable == "platano" || variable == "fresa":
		fmt.Println("Es una fruta")		
	default:
		fmt.Println("no es un animal ni fruta")
	}
}