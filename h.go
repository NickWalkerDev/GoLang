package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func readers() {
	reads, err := os.Create("test.txt")

	if err != nil {
		log.Fatal(err)
	}

	defer reads.Close()

	scanner := bufio.NewScanner(reads)

	for scanner.Scan() {

		fmt.Println(scanner.Text())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	text := "Text overwrite of file "
	overWrite := []byte(text)

	_, err1 := reads.Write(overWrite)

	if err1 != nil {
		log.Fatal(err1)
	}

}

func main() {
	readers()

}
