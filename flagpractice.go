package main

import (
	"flag"
	"fmt"
)

func main() {
	var limitFlag = flag.Int("limit", 1000, "This is the limit of lines per file")
	flag.Parse()
	fmt.Println(*limitFlag)
}
