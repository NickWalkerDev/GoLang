package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
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

	fmt.Println("Updated file.")
}

// reads file in console. Error message if file doesn't open.
func readFiles() {
	file, err := os.Open("thisisatest.txt")

	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var txtlines []string

	for scanner.Scan() {
		txtlines = append(txtlines, scanner.Text())
	}

	file.Close()

	for _, eachline := range txtlines {
		fmt.Println(eachline)
	}
}

// temporarily rename the file
func renameFiles() {
	oldName := "thisisatest.txt"
	newName := "opened-texted-file.txt"
	err := os.Rename(oldName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

// rever updated name of file back to original name of file.
func revertName() {
	currentName := "opened-texted-file.txt"
	newName := "thisisatest.txt"
	time.Sleep(7 * time.Second)
	err := os.Rename(currentName, newName)
	if err != nil {
		log.Fatal(err)
	}
}

// print the words to the terminal of file.
func wordScanner() {
	// Open a file.
	f, _ := os.Open("thisisatest.txt")

	reader := bufio.NewReader(f)
	content, _ := ioutil.ReadAll(reader)
	fmt.Println("The document contained the text:")
	// File content.
	fmt.Println(string(content))

}

// move the file from VS Code to desktop in untitled-folder.
func moveFile() {
	oldLocation := "thisisatest.txt"
	newLocation := "/Users/nickwalker/Desktop/untitled-folder"
	err := os.Rename(oldLocation, newLocation)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	createFile()
	writeFile()
	readFiles()
	renameFiles()
	revertName()
	wordScanner()
	moveFile()

}
