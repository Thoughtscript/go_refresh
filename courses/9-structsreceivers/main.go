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
	cow := Animal{"cow", "grass", "walk", "moo"}
	bird := Animal{"bird", "worms", "fly", "peep"}
	snake := Animal{"snake", "mice", "slither" ,"hsss"}

	fmt.Println(">")
	bscan := bufio.NewScanner(os.Stdin)

	for bscan.Scan() {
		line := bscan.Text()
		if line == "" {
			fmt.Println("Please select one of [cow, bird, snake] and [eat, move, speak] separated by a space...")
			fmt.Println("e.g. - 'cow move' \n")
		} else {
			arr := strings.Split(line, " ")

			switch arr[0] {
				case "cow":
					switch arr[1] {
						case "move": cow.Move()
						case "eat": cow.Eat()
						case "speak": cow.Speak()
						default: fmt.Println("Please select one of [eat, move, speak] \n")
					}
				case "bird":
					switch arr[1] {
						case "move": bird.Move()
						case "eat": bird.Eat()
						case "speak": bird.Speak()
						default: fmt.Println("Please select one of [eat, move, speak] \n")
					}
				case "snake":
					switch arr[1] {
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