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

var data = map[string]Animal{
	"cow": {
		food:       "grass",
		locomotion: "walk",
		noise:      "moo",
	},
	"bird": {
		food:       "worms",
		locomotion: "fly",
		noise:      "peep",
	},
	"snake": {
		food:       "mice",
		locomotion: "slither",
		noise:      "hsss",
	},
}

func (a Animal) Eat() {
	fmt.Println(a.food)
}

func (a Animal) Move() {
	fmt.Println(a.locomotion)
}

func (a Animal) Speak() {
	fmt.Println(a.noise)
}

func main() {
	var animal Animal
	for {
		reader := bufio.NewReader(os.Stdin)
		fmt.Print(">")
		request, _ := reader.ReadString('\n')
		words := strings.Fields(request)
		if len(words) != 2 {
			println("needs exactly two words")
			continue
		}
		switch words[0] {
		case "cow":
			animal = data["cow"]
		case "bird":
			animal = data["bird"]
		case "snake":
			animal = data["snake"]
		default:
			println("unknown animal")
		}

		switch words[1] {
		case "speak":
			animal.Speak()
		case "move":
			animal.Move()
		case "eat":
			animal.Eat()
		default:
			println("unknown action")
		}
	}
}
