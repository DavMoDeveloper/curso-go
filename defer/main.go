package main

import (
	"os"
	"fmt"
)

func main() {
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Printf("Ocurrió un error al crear el archivo %v", err)
		return
	}
	defer file.Close()
	
	_, err = file.Write([]byte("Hello Gopers"))
	if err != nil {
		fmt.Printf("Ocurrio un error al escribir el archivo: %v", err)
		return
	}

	file.Close()
}