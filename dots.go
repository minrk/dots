package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// read each line from stdin...
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// and write a dot to stdout
		fmt.Print(".")
	}
	fmt.Println()
}
