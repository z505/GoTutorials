// Shows how to obtain arguments from a command line program...
// i.e. at the command if you run 
//   somecommand -abc -123

package main

import (
	"fmt"
	"os"
)

func showArg1() {
	fmt.Print(`Arg 1: `)
	if len(os.Args) > 1 {
		fmt.Print(os.Args[1])
	}
	fmt.Println()
}

func showArgs() {
	fmt.Println(`How many arguments: `, len(os.Args)-1)
	for i, arg := range os.Args[1:] {
		fmt.Println(`Arg `, i+1, `: `, arg)
	}
}

func main() {
	fmt.Println(`Example to show command arguments: `)
	showArg1()
	showArgs()
}
