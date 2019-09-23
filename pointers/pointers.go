package main

import "fmt"

func zeroval(ival int) {
	ival = 0
}

func zeroptr(iptr *int) {
	*iptr = 0
}

func main() {
	i := 1
	fmt.Println("Init:", i)
	
	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	// Overwrites the pre-main functions because it points to memory. cool.
	fmt.Println("zeroptr:", i)

	fmt.Println("pointers:", &i)
}