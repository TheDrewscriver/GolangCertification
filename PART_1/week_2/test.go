package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	var line string
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		line = scanner.Text()
	}
	fmt.Printf("Input was: %q\n", line)
}
