package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

var path = "thisisatest.txt"

func createFile() {
	// check if file exists
	var _, err = os.Stat(path)

	// create file if not exists
	if os.IsNotExist(err) {
		var file, err = os.Create(path)
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()
	}

	fmt.Println("Created file", path)
}

// give timestamp in console of when a change was updated. This also writes to file when the change was made.
func writeFile() {
	timeNow := time.Now()
	fmt.Println(timeNow.String())

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// Write some text line-by-line to file.
	_, err = file.WriteString("A change was made at: \n" + time.Now().String())
	if err != nil {
		log.Fatal(err)
	}

	// Save file changes.
	err = file.Sync()
	if err != nil {
		log.Fatal(err)
	}

}

func main() {
	// ./flagpractice -from test.txt -to /home/nick/downloads/test.txt
	fromFlag := flag.String("from", "", "from file")
	toFlag := flag.String("to", "", "to file")
	flag.Parse()
	if *fromFlag == "" {
		log.Fatalf("Specifiy -from file")
	}
	if *toFlag == "" {
		log.Fatalf("Specify -to file")
	}
	log.Printf("fromFlag: %v; toFlag: %v", *fromFlag, *toFlag)
	err := os.Rename(*fromFlag, *toFlag)
	if err != nil {
		log.Fatal(err)
	}

	// createFile()
	// writeFile()

	// addcmd := flag.NewFlagSet("add", flag.ExitOnError)
	// a_add := addcmd.Int("a", 0, "The value of a")
	// b_add := addcmd.Int("b", 5, "The value of b")

	// mulcmd := flag.NewFlagSet("mul", flag.ExitOnError)
	// a_mul := mulcmd.Int("a", 147, "The value of a")
	// b_mul := mulcmd.Int("b", 234, "The value of b")
	// flag.Parse()

	// switch os.Args[1] {
	// case "add":
	// 	addcmd.Parse(os.Args[2:])
	// 	fmt.Println(*a_add + *b_add)
	// case "mul":
	// 	mulcmd.Parse(os.Args[2:])
	// 	fmt.Println(*(a_mul) * (*b_mul))
	// default:
	// 	fmt.Println("expected 'add' or 'mul' command")
	// 	os.Exit(1)
	// }

}
