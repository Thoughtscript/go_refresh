package main

import "fmt"

func Swap(arr []int, x int, y int) {
	temp := arr[x]
	arr[x] = arr[y]
	arr[y] = temp
}

func BubbleSort(nums int) {

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

	for i := 0; i < len(arr) -1; {
		if arr[i] > arr[i+1] {
			Swap(arr, i, i+1)
			fmt.Printf("Swapped num %d at index %d with num %d at index %d\n", arr[i], i,  arr[i+1], i+1)
			fmt.Printf("Array state: %d \n", arr)
			i = 0
		} else {
			i++
		}
	}

}

func main() {
	BubbleSort(10)
}
