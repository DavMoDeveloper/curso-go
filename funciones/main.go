package main

import "fmt"

func main() {
	emoji := "Perro"
	change(&emoji)
	fmt.Println(emoji)	
	//hello("David", "Moreno")
}
func change(value *string) {
	*value = "Hola"
}



/*func hello() {
	fmt.Println("Hello EDTeam")
}*/

//Funciones con parametro
/*func hello(firstName string, lastName string) {
	fmt.Printf("Hello %s %s", firstName, lastName)
}*/