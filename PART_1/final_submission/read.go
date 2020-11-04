package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

//The struct we are storing data in
type Person struct {
	fName string
	lName string
}

/*
 Write a program which reads information from a file and represents it in a slice of structs.
Assume that there is a text file which contains a series of names.
Each line of the text file has a first name and a last name, in that order, separated by a single space on the line.
Your program will define a name struct which has two fields, fname for the first name, and lname for the last name.
Each field will be a string of size 20 (characters).
Your program should prompt the user for the name of the text file.
Your program will successively read each line of the text file and create a struct which contains the first and last names found in the file.
Each struct created will be added to a slice, and after all lines have been read from the file, your program will have a slice containing one struct for each line in the file.
After reading all lines from the file, your program should iterate through your slice of structs and print the first and last names found in each struct.
Submit your source code for the program, “read.go”.
*/
func main() {
	//Variable to store fileName
	var fileName string
	//Make the slice to store the objects
	slice := make([]Person, 0, 1)
	//Read the filename
	fmt.Scanln(&fileName)
	//Open
	fd, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
	}
	//Create scanner
	scanner := bufio.NewScanner(fd)
	//For each line
	for scanner.Scan() {
		//Read text
		input := scanner.Text()
		//Split on space
		splitString := strings.Split(input, " ")
		//Extract firstname and lastname
		firstName := splitString[0]
		lastName := splitString[1]
		//Instantiate the object
		personObject := Person{fName: firstName, lName: lastName}
		//Add to slice
		slice = append(slice, personObject)
	}
	//Print each element to slice
	for _, x := range slice {
		fmt.Printf("%s %s \n", x.fName, x.lName)
	}

}
