package main
import (
	"fmt"

	"github.com/DavMoDeveloper/curso-go/paquetes/greet"
	"rsc.io/quote/v3"
)

func main() {
	fmt.Println(greet.Italian())
	fmt.Println(quote.HelloV3())
	fmt.Println(quote.Concurrency())
}