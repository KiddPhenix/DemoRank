// In order to make the log more detailed in the future, Encapsulate it first
package toolkit

import (
	//"bytes"
	//"fmt"
	"log"
)

func L(v ...any) {
	log.Print(v)
}

func F(v ...any) {
	log.Fatal(v)
}
