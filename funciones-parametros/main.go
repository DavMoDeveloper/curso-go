package main

import (
	"fmt"
)

func main() {
	/*nums := []int{1,4,70,5,67,90,2}
	results := filter(nums, func(num int) bool{
		if num <= 10 {
			return true
		}
		return false
	})
	fmt.Println(results)*/
	x := hello("David")("Moreno")
	fmt.Println(x)
}

func hello(name string) func(string) string {
	return func(apellido string) string {
		return "Hello " + name + " " + apellido
	}
}

func filter(nums []int, callback func(int) bool) []int {
	result := []int{}
	for _, v := range nums {
		if callback(v){
			result = append(result, v)
		}
	}

	return result
}