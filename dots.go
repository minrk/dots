package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	dot := "."
	if len(os.Args) > 1 {
		// `dots ARG` uses ARG instead of '.'
		dot = os.Args[1]
	}

	// read each line from stdin...
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		// and write a dot to stdout
		fmt.Print(dot)
	}
	fmt.Println()
}
