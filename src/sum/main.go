package main

/*
int sum(int a, int b) {
	return a + b;
}
*/
import "C"
import (
	"fmt"
	"math/rand/v2"
)

func main() {
	a := rand.Int32N(1000)
	b := rand.Int32N(1000)

	s := C.sum(C.int(a), C.int(b))
	fmt.Printf("%d + %d = %d\n", a, b, s)
}
