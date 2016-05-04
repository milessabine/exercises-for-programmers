/*What is the input string? Homer
Homer has 5 characters.*/
package main

import "fmt"

func main() {
	//read input
	//n:= getInput("What is the input string?")
	fmt.Printf("What is the input string? ")
	var n string
	fmt.Scanln(&n)

	//figure out how many characters input string has
	c := len(n)
	if c == 0 {
		fmt.Println("You must enter something.")
		return
	}

	//printing program output
	//o := n + " has " + c + " characters."
	fmt.Printf("%s has %d characters.\n", n, c)
}
