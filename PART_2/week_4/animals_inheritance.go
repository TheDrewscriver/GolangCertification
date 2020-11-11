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

//Snake implements interface
type Snake struct{ name string }

func (s Snake) Eat() {
	fmt.Println("mice")
}

func (s Snake) Move() {
	fmt.Println("slither")
}

func (s Snake) Speak() {
	fmt.Println("hsss")
}

//Bird implements interface
type Bird struct{ name string }

func (b Bird) Eat() {
	fmt.Println("worms")
}

func (b Bird) Move() {
	fmt.Println("fly")
}

func (b Bird) Speak() {
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
			continue
		}

		fmt.Println(splitValue)
		switch splitValue[0] {

		case "newanimal":
			{
				nameValue := splitValue[1]
				typeValue := splitValue[2]
				switch typeValue {
				case "cow":
					hashMap[nameValue] = Cow{nameValue}
				case "snake":
					hashMap[nameValue] = Snake{nameValue}
				case "bird":
					hashMap[nameValue] = Bird{nameValue}
				default:
					fmt.Println("Only bird,snake and cow are supported as animals")
				}
			}
		case "query":
			{
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
		default:
			fmt.Println("Only newanimal and query are supported as options")
		}

		// var a Animal
		// var c Cow= Cow{"Brian"}
		// a=c
		// hashMap["Brian"]=a
		// hashMap["Brian"].Eat()

	}
}
