package main

import (
	"fmt"
	"encoding/json"
	"bufio"
	"os"
)

func main() {
	/*
		Write a program which prompts the user to first enter a name, and then enter an address.
		Your program should create a map and add the name and address to the map using the keys “name” and “address”, respectively.
		Your program should use Marshal() to create a JSON object from the map, and then your program should print the JSON object.
	*/
	//Make a map
	var person=make(map[string]string)
	//Get name
	input := readInput()
	person["name"] = input
	//Get address
	input = readInput()
	person["address"] = input
	//Marshall JSON
	output,err:=json.Marshal(person)
	if(err!=nil){
		fmt.Println("Could not marshal")	
		return
	}
	fmt.Println(string(output))
}

/*Function to read input from the command line,
using bufio to accept spaces*/
func readInput() string{
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}
