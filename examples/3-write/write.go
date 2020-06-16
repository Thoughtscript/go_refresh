package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	toWrite := []byte("Hello World!")

	// If you execute this in the main root it'll appear two directories up!
	dErr := ioutil.WriteFile("output.txt", toWrite, 0777)

	// Error handling
	if dErr != nil {
		fmt.Println("Exception encountered " + dErr.Error())
	}

}
