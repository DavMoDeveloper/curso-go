package main
import (
	"fmt"
	"strings"
)

func main() {
	texto := "DaViD MoREno"
	minusc, mayusc := convert(texto)
	fmt.Println(minusc, mayusc)
	
	//total := sum(2, 3)
	//println(total)
}

func sum(num1, num2 int) int {
	return num1 + num2
}

//retorno de multiples valores
func convert(text string) (string, string) {
	min := strings.ToLower(text)
	may := strings.ToUpper(text)
	return min, may
}