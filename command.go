package main

import (
	"flag"
	"fmt"
	"os"
)

// allows flags in command line
var flagvar int

func init() {
	flag.IntVar(&flagvar, "command.go", 0, "help output for command.go")
	flag.Parse()

}

//creates a new txt file then multiplies each number in there by two and outputs it in same txt file.
func multiplyValues() {
	f, err := os.Create("multiplyValues.txt")
	if err != nil {
		fmt.Printf("error creating file: %v", err)
		return
	}

	defer f.Close()
	for i := 0; i < 16; i++ {
		f.WriteString("Multiply by two: ")
		_, err = f.WriteString(fmt.Sprintf("%d, ", i*2))
		_, err = f.WriteString(fmt.Sprintf("%d\n\n ", i))
		if err != nil {
			fmt.Printf("error writing string: %v", err)
		}
	}

}

func main() {
	// displays the name of the program in terminal
	cmd := os.Args[0]
	fmt.Printf("Program Name: %s\n", cmd)

	multiplyValues()
}
