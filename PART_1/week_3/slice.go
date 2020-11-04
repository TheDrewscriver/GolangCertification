package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	/*Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
	The program should be written as a loop.
	Before entering the loop, the program should create an empty integer slice of size (length) 3.
	During each pass through the loop, the program prompts the user to enter an integer to be added to the slice.
	The program adds the integer to the slice, sorts the slice, and prints the contents of the slice in sorted order.
	The slice must grow in size to accommodate any number of integers which the user decides to enter.
	The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
	*/
	//This stores the input we read
	var input string
	//Instantiate array of size 3 as required
	var array [3]int
	//Create a slice to store our input integers
	slice := array[1:1]
	//Loop
	for true {
		//Read input
		fmt.Scanln(&input)
		//If input is X or x, leave the loop
		if input == "X" || input == "x" {
			break
		}
		//Else parse the integer. If not an integer, we ignore
		integerInput, err := strconv.Atoi(input)
		if err != nil {
			continue
		}
		//Append to slice
		slice = append(slice, integerInput)
		//Sort the slice
		sort.Ints(slice)
		//Print sorted slice
		fmt.Println(slice)
	}
}
