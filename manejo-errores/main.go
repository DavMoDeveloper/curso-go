package main

import (
	"fmt"
	//"io/ioutil"
	"errors"
)

func main() {
	/*content , err := ioutil.ReadFile("./things.txt")
	if err != nil {
		fmt.Println("Ocurrio un error: &v", err)
		return
	} 	
	
	fmt.Println(string(content))*/
	result, err := division(10, 2)
	if err != nil {
		fmt.Println("Ocurrio un error: &v", err)
		return
	} 	
	
	fmt.Println(result)
}

func division(dividendo, divisor int) (int, error) {
	if divisor == 0 {
		return 0, errors.New("No puedes dividir por cero")
	}

	return dividendo / divisor, nil
}
