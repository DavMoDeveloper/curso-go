package main

import (
	"fmt"
)

func main() {
	//For clasico
	/*for i := 1; i <= 10 ; i++ {
		fmt.Println(i)
	}*/

	//For continuo
	/*i := 1;
	for i <= 10  {
		fmt.Println(i)
		i++
	}*/

	//For ever
	/*i := 1;
	for {
		fmt.Println(i)
		i++
	}*/

	//For rage
	/*nums:= []uint8{2,4,6,8}

	for i, v := range nums {
		fmt.Println("Indice", i, "Valor", v)
		nums[i] *=2
	}
	fmt.Println(nums)*/

	//Iterar sobre un map
	/*sports := map[string]string{"basketball":"baloncesto", "soccer":"futbol"}
	
	for key, v := range sports {
		fmt.Println("Key", key, "Valor", v)
	}*/
	
	//Iterar string
	hello := "hello"
	for _, v := range hello {
		fmt.Println(string(v))
	} 
	
}