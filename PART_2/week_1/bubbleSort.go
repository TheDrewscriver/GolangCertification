package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("Enter integers separated by space, and press [ENTER]")
	input := readInput()
	//fmt.Println(input)
	array := convertInputStringToIntArray(input)
	//fmt.Printf("Input parsed %d", array)
	bubbleSort(array)
	fmt.Println(array)
}

func convertInputStringToIntArray(inputString string) []int {
	//Split string by space
	tmp := strings.Split(inputString, " ")
	values := make([]int, 0)
	for _, value := range tmp {
		v, err := strconv.Atoi(value)
		if err != nil {
			continue
		}
		values = append(values, v)
	}
	return values
}

func bubbleSort(array []int) {
	//Get size of the array
	n := len(array)
	//For the length of the array
	for i := 0; i < n-1; i++ {
		//For the length of array minus sorted elements
		for j := 0; j < (n - i - 1); j++ {
			//If the current element is greater than the next
			if array[j] > array[j+1] {
				//Swap
				swap(array, j)
			}
		}
	}
}

func swap(slice []int, i int) {
	//Store current element in a temp variable
	temp := slice[i]
	//Replace current element with next
	slice[i] = slice[i+1]
	//Replace next element with the temp variable
	slice[i+1] = temp
}

/*Function to read input from the command line,
using bufio to accept spaces*/
func readInput() string {
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		input = scanner.Text()
	}
	return input
}
