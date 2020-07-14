package main

import (
	"fmt"
	"time"
)

/**
 * Concurrent, asynchronous code example - e.g. race conditions.
 *
 * Code that depends on two asynchronous but concurrent tasks being executed.
 */

func b(z int, id int) int {
	if z > 3 { time.Sleep(1000) }
	fmt.Println(z, id)
	return z
}

func a(x int, y int, id int) int {
	if x < 5 {
		time.Sleep(100)
		fmt.Println(x, id)
		return x
	} else if x + y > 10 {
		time.Sleep(1000)
		fmt.Println(x+y, id)
		return x + y
	}

	fmt.Println(y, id)
	return b(y, id)
}

func main() {
	nums := [][]int{{1,1},{0,0},{2,1},{100,200},{15,244}}
	/*
		Go routines that are literally placed sequentially will not print to console
		They must either print out to a channel or be in a loop
	*/

	/* You'll notice that a couple of the operations will vary based on input (e.g. - "race conditions"): */

	/*
	1 0
	0 1
	2 2
	2 2
	300 3
	1 0
	300 3
	0 1
	259 4
	*/

	/*
	1 0
	0 1
	2 2
	2 2
	300 3
	0 1
	300 3
	259 4
	1 0
	*/

	for i := 0; i < len(nums); i++ {
		// Two functions with race conditions
		go b(a(nums[i][0], nums[i][1], i), i)
	}
}