package main

import (
	"flag"
	"fmt"
	"log"
	"os"
)

func main() {

	fileFlag := flag.String("file", "", "Please name it ending in txt")
	contentFlag := flag.String("content", "", "Write contents of file")
	flag.Parse()
	f, err := os.Create(*fileFlag)
	if *fileFlag == "" {
		log.Fatalf("Specify the name of the file -file")
		fmt.Println(err)
	}

	_, err = f.WriteString(*contentFlag)
	if *contentFlag == "" {
		log.Fatalf("Specify the context -content")

	}

	// command line: ./writeFiles -create "nick.txt" -write "went to the store today"
}
