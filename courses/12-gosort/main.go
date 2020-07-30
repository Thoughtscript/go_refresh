package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

func buildLastArr(a []int, b []int, c []int, d []int) []int {
	finalArr := append(a, b...)
	finalArr = append(finalArr, c...)
	finalArr = append(finalArr, d...)
	return finalArr
}

func finalSort(finalArr []int) {
	fmt.Printf("Assembling array: %v \n", finalArr)
	for i := 0; i < len(finalArr) - 1; {
		if finalArr[i] > finalArr[i+1] {
			orig := finalArr[i]
			finalArr[i] = finalArr[i+1]
			finalArr[i+1] = orig
			i = 0
		} else {
			i++
		}
	}
	fmt.Printf("Final sort: %v - length: %d \n", finalArr, len(finalArr))
}

func sortAsync(wg *sync.WaitGroup, arr []int, c chan []int) {
	defer wg.Done()
	fmt.Printf("Go begin: %v \n", arr)
	for i := 0; i < len(arr) - 1; {
		if arr[i] > arr[i+1] {
			orig := arr[i]
			arr[i] = arr[i+1]
			arr[i+1] = orig
			i = 0
		} else {
			i++
		}
	}
	c <- arr
	fmt.Printf("Go end: %v \n", arr)
}

func main() {
	// Create necessary resources
	wg := new(sync.WaitGroup)
	c := make(chan []int)
	arrs := make([][]int, 4)

	// User input of arrays
	fmt.Println("Enter a series of four or more int delimited by commas...")
	fmt.Println("e.g. - 1,2,4,5,6,9,9,9213,24,4,5,6")
	bscan := bufio.NewScanner(os.Stdin)

	for bscan.Scan() {
		line := bscan.Text()

		// Exit
		if line == "exit" {
			os.Exit(0)
		}

		strArr := strings.Split(line, ",")
		fmt.Printf("You entered: %v - length: %d\n", strArr, len(strArr))
		if len(strArr) < 4 {
			fmt.Printf("Please enter four or more int delimited by commas... \n")
			os.Exit(0)
		}
		counter := 0
		for i := 0; i < len(strArr); i++ {
			// These conversion error checks must be here!
			j, err := strconv.Atoi(strArr[i])
			if err != nil {
				fmt.Printf("Please enter only int delimited by commas... \n")
				os.Exit(0)
			}
			arrs[counter] = append(arrs[counter], j)
			counter++
			if counter > 3 {
				counter = 0
			}
		}

		fmt.Printf("Selected arrays: %v \n", arrs)

		for i := 0; i < len(arrs); i++ {
			wg.Add(1)
			fmt.Printf("Sorting: %v \n", arrs[i])
			go sortAsync(wg, arrs[i], c)
		}

		// Wait for each completed sort here
		// If a channel is not used and a sort is attempted - it will cause a deadlock issue
		sOne := <- c
		sTwo := <- c
		sThree := <- c
		sFour := <- c

		time.Sleep(10000)

		finalSort(buildLastArr(sOne, sTwo, sThree, sFour))
		wg.Wait()
	}
}