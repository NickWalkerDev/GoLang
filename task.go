//read file takes each line
// parse each line as a number * 2
// print to outut file if no output selected it justs prints

// -go build task.go && ./task -parse "text.txt" goes through each line of txt file and multiplies it by 2
// go build task.go && ./task -print "text.txt" select which file to print to or print to console if none selected

package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
)

func linesInFile(fileName string) ([]string, error) {
	f, err := os.Open(fileName)
	if err != nil {
		return nil, err
	}
	// Create new Scanner.
	scanner := bufio.NewScanner(f)
	result := []string{}
	// Use Scan.
	for scanner.Scan() {
		line := scanner.Text()
		result = append(result, line)
	}
	return result, nil
}

func main() {
	parseFlag := flag.String("parse", "", "-parse to go through a txt file and multiply by two")
	writeFlag := flag.String("write", "", "-write writes to new file the parsed data")
	flag.Parse()
	// lines := linesInFile(*parseFlag)
	// for index, line := range  {
	// 	fmt.Printf("Index = %v, line = %v\n", index, line)
	// }

	// Get count of lines.
	lines, err := linesInFile(*parseFlag)
	x := strconv.Itoa(len(lines) * 2)
	if err != nil {
		log.Fatal(err)
	}

	f, err := os.Create(*writeFlag)
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(x)
	if err != nil {
		fmt.Println(err, l)
		f.Close()
		return
	}

}
