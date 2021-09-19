package main

import (
	"fmt"
)

func main() {
	var dog, cat string = "Camote", "Michi"
	fmt.Println(dog, cat)
	var dog1, cat2 = "Fido", "Minino"
	fmt.Println(dog1,cat2)
	dog3, cat3 := "Asan", "Michi"
	fmt.Println(dog3,cat3)
}