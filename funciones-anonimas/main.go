package main

import "fmt"

func main() {
	/*x := func() {
		fmt.Println("Hello Edteam")
	}
	x()*/

	//Autoejecutada
	func(text string) {
		fmt.Println("Hello ", text)
	}("Gopers")
}