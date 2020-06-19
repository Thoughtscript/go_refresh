package main

import "fmt"

func Swap(slice []int, i int) {
	temp := slice[i]
	slice[i] = slice[i+1]
	slice[i+1] = temp
}

func BubbleSort(slice []int) {

	for i := 0; i < len(slice) -1; {
		if slice[i] > slice[i+1] {
			Swap(slice[0 : len(slice)], i)
			fmt.Printf("Swapped num %d at index %d with num %d at index %d\n", slice[i], i,  slice[i+1], i+1)
			fmt.Printf("Array state: %d \n", slice)
			i = 0
		} else {
			i++
		}
	}

}

func main() {
	
	nums := 10
	// array
	arr := make([]int, nums, nums)

	for i := 0; i < nums; i++ {
		num := 0
		remaining := nums - i

		// Enter an int
		fmt.Printf("Enter an integer (%d  more to enter): \n", remaining)
		fmt.Scan(&num)
		arr[i] = num
		fmt.Printf("Numbers entered: %d \n", arr)
	}

	fmt.Println("\n Sorting array... \n")

	BubbleSort(arr[0:len(arr)])
}
