// HelloWorld project main.go
package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		os.Exit(1)
	}

	var eu string
	eu = "roberto"
	name, power := 0, "outro roberto"
	fmt.Println("Hello World!", eu, os.Args[1], name, power)
	log("print from func")
	hasAdded, newEu := givePower(eu)
	if hasAdded {
		fmt.Println("hasAdded? ", hasAdded, newEu)
	}

	hasAdded, outro := givePower("outro")
	if !hasAdded {
		fmt.Println("hasAdded? ", hasAdded, outro)
	}
}

func log(message string) {
	fmt.Println(message)
}

func add(a int, b int) int {
	return 0
}

func givePower(name string) (bool, string) {
	b := false
	if name == "roberto" {
		name += " added"
		b = true
	} else {
		name += " se fudeu"
	}
	return b, name
}
