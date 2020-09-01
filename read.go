package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	readFlag := flag.String("read", "", "-read to see flag")
	flag.Parse()
	data, err := ioutil.ReadFile(*readFlag)
	if err != nil {
		log.Panicf("Enter a file name to read")
	}
	text := string(data)
	fmt.Println(text)

}
