package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//The interface
type Animal interface {
	Eat()
	Move()
	Speak()
}

//Cow implements interface
type Cow struct{ name string }

func (c Cow) Eat() {
	fmt.Println("grass")
}

func (c Cow) Move() {
	fmt.Println("walk")
}

func (c Cow) Speak() {
	fmt.Println("moo")
}

//Bird implements interface
type Bird struct{ name string }

func (b Bird) Eat() {
	fmt.Println("mice")
}

func (b Bird) Move() {
	fmt.Println("slither")
}

func (b Bird) Speak() {
	fmt.Println("hsss")
}

//Snake implements interface
type Snake struct{ name string }

func (s Snake) Eat() {
	fmt.Println("worms")
}

func (s Snake) Move() {
	fmt.Println("fly")
}

func (s Snake) Speak() {
	fmt.Println("peep")
}

func main() {
	/*
		The variables we use
	*/

	hashMap := make(map[string]Animal)
	var input string
	for true {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)
		if scanner.Scan() {
			input = scanner.Text()
		}

		splitValue := strings.Split(input, " ")
		if len(splitValue) != 3 {
			fmt.Println("The input string should have 3 inputs")
			os.Exit(1)

		}

		fmt.Println(splitValue)
		switch splitValue[0] {

		case "newanimal":
			{
				fmt.Println("inputValue " + splitValue[0])
				nameValue := splitValue[1]
				fmt.Println(nameValue)

				typeValue := splitValue[2]
				fmt.Println(typeValue)
				switch typeValue {
				case "cow":
					hashMap[nameValue] = Cow{nameValue}
					fmt.Println("case cow")
					fmt.Println(hashMap)
				case "snake":
					hashMap[nameValue] = Snake{nameValue}
				case "bird":
					hashMap[nameValue] = Bird{nameValue}

				}

			}
		case "query":
			{
				fmt.Println("Here")
				nameValue := splitValue[1]
				objectValue := hashMap[nameValue]
				if objectValue == nil {
					fmt.Println("Unable to find the animal you queried")
					continue
				}

				function := splitValue[2]
				switch function {
				case "move":
					hashMap[nameValue].Move()
				case "eat":
					hashMap[nameValue].Eat()
				case "speak":
					hashMap[nameValue].Speak()
				default:
					fmt.Println("Unrecognized animal action")
				}

			}
		}

		// var a Animal
		// var c Cow= Cow{"Brian"}
		// a=c
		// hashMap["Brian"]=a
		// hashMap["Brian"].Eat()

	}
}
