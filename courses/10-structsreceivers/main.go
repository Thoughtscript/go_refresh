package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	animalName string
	foodEaten string
	locomotionMethod string
	spokenSound string
}

func (a Animal) Eat() {
	fmt.Println(a.foodEaten + "\n")
}

func (a Animal) Move() {
	fmt.Println(a.locomotionMethod + "\n")
}

func (a Animal) Speak() {
	fmt.Println(a.spokenSound + "\n")
}

func main() {

	var animals = make(map[string]string)

	cow := Animal{"cow", "grass", "walk", "moo"}
	bird := Animal{"bird", "worms", "fly", "peep"}
	snake := Animal{"snake", "mice", "slither" ,"hsss"}

	fmt.Println(">")
	bscan := bufio.NewScanner(os.Stdin)

	for bscan.Scan() {
		line := bscan.Text()
		arr := strings.Split(line, " ")

		if arr[0] == ""  {
			fmt.Println("Enter 'query' to print an animal or enter 'newanimal' to create a new one")
			fmt.Println("e.g. - > query cow move")
			fmt.Println("e.g. - > newanimal sally cow \n")

		} else if arr[0] == "newanimal" {
			switch arr[2] {
				case "cow":
					animals[arr[1]] = "cow"
					break
				case "bird":
					animals[arr[1]] = "bird"
					break
				case "snake":
					animals[arr[1]] = "snake"
					break
				default:
					fmt.Println("Please select a name and a type [cow, bird, snake] \n")
			}
			fmt.Println("Created it! \n")

		} else if arr[0] == "query" {

			if !(arr[1] == "cow" || arr[1] == "bird" || arr[1] == "snake") {
				name := arr[1]
				arr[1] = animals[name]
			}

			switch arr[1] {
				case "cow":
					switch arr[2] {
						case "move": cow.Move()
						case "eat": cow.Eat()
						case "speak": cow.Speak()
						default: fmt.Println("Please select one of [eat, move, speak] \n")
					}
				case "bird":
					switch arr[2] {
						case "move": bird.Move()
						case "eat": bird.Eat()
						case "speak": bird.Speak()
						default: fmt.Println("Please select one of [eat, move, speak] \n")
					}
				case "snake":
					switch arr[2] {
						case "move": snake.Move()
						case "eat": snake.Eat()
						case "speak": snake.Speak()
						default: fmt.Println("Please select one of [eat, move, speak] \n")
					}
				default: fmt.Println("Please select one of [cow, bird, snake] \n")
			}
		}
	}
}