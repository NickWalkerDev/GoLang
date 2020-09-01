package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	addcmd := flag.NewFlagSet("add", flag.ExitOnError)
	a_add := addcmd.Int("a", 0, "The value of a")
	b_add := addcmd.Int("b", 5, "The value of b")

	mulcmd := flag.NewFlagSet("mul", flag.ExitOnError)
	a_mul := mulcmd.Int("a", 147, "The value of a")
	b_mul := mulcmd.Int("b", 234, "The value of b")
	flag.Parse()

	switch os.Args[1] {
	case "add":
		addcmd.Parse(os.Args[2:])
		fmt.Println(*a_add + *b_add)
	case "mul":
		mulcmd.Parse(os.Args[2:])
		fmt.Println(*(a_mul) * (*b_mul))
	default:
		fmt.Println("expected 'add' or 'mul' command")
		os.Exit(1)
	}

}
