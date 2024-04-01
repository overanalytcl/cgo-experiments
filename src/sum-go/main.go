package main

//#include "sumgo.h"
import "C"
import (
	"math/rand/v2"
)

//export sum
func sum(a, b C.int) C.int {
	return a + b
}

func main() {
	a := rand.Int32N(1000)
	b := rand.Int32N(1000)

	C.print_sum(C.int(a), C.int(b))
}
