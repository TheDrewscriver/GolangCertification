package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Animal struct {
	food       string
	locomotion string
	noise      string
}

func (a *Animal) InitMe(food string, locomotion string, noise string) {
	a.food = food
	a.locomotion = locomotion
	a.noise = noise
}

func (a *Animal) Eat() {
	fmt.Println(a.food)
}

func (a *Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a *Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	/*
		The variables we use
	*/

	cow := new(Animal)
	cow.InitMe("grass", "walk", "moo")

	bird := new(Animal)
	bird.InitMe("worms", "fly", "peep")

	snake := new(Animal)
	snake.InitMe("mice", "slither", "hsss")

	var input string

	for true {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}

		splitValue := strings.Split(input, " ")
		if strings.EqualFold(splitValue[0], "cow") {
			executeFunction(*cow, splitValue[1])
		}
		if strings.EqualFold(splitValue[0], "bird") {
			executeFunction(*bird, splitValue[1])
		}
		if strings.EqualFold(splitValue[0], "snake") {
			executeFunction(*snake, splitValue[1])
		}
	}
}

func executeFunction(animal Animal, value string) {

	if strings.EqualFold(value, "eat") {
		animal.Eat()
	}
	if strings.EqualFold(value, "move") {
		animal.Move()
	}
	if strings.EqualFold(value, "speak") {
		animal.Speak()
	}

}
