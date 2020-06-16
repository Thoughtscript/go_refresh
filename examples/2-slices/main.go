package main
// Note that the main method must be in the package main

import "fmt"

func main() {
	x := [...]int {1, 2, 3, 4, 5}
	fmt.Println(x, len(x), cap(x))
	//[1 2 3 4 5] 5 5

	//[n:m] is inclusive - exclusive - same almost all substr and slice operations in any language!
	//capacity is the calculated by taking the start index of the slice to the end of the sliced index
	y := x[0:2]
	fmt.Println(y, len(y), cap(y))
	//[1 2] 2 5
	//capacity is 5 -> 1,2,[3,4,5]

	z := x[1:4]
	fmt.Println(z, len(z), cap(z))
	//[2 3 4] 3 4
	//capacity is 4 -> [],2,3,4,5
}