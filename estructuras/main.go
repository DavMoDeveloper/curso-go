package main

import (
	"fmt"
)

func main() {
	type course struct {
		Name string
		Professor string
		Country string
	}

	db := course{
		Name: "Bade de datos",
		Professor: "David",
		Country: "Perú",
	}
	fmt.Printf("%+v\n", db)

	/*git := course{"Git","Beto","Bolivia"}
	fmt.Printf("%+v\n", git)

	css := course{Professor: "Alvaro"}
	fmt.Printf("%+v\n", css)*/

	p := &db
	(*p).Professor = "Beto"
	p.Name = "Android"
	fmt.Printf("%+v\n", db)
	fmt.Printf("%+v\n", p)
}