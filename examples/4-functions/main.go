package main

import ("fmt")

// Higher-order function
// Note how the entire function signature is passed as a parameter
func applyIt(afunct func (int) int, val int) int {
	return afunct(val)
}

// Higher-order function
// Return a function
func returnFunction() func (x int) int {
	return func (x int) int {
		return x
	}
}

// name, args, return
func incFn(x int) int {return x + 1}
func decFn(x int) int {return x - 1}

func main() {
	fmt.Println(applyIt(incFn, 2))
	fmt.Println(applyIt(decFn, 21))
	fmt.Println(returnFunction()(999))

	// anonymous functions
	fmt.Println(applyIt(func (x int) int {return x * 2}, 2))
}