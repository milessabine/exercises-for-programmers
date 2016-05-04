/*
What is the quote? These arn't the droids you're looking for.
Who said it? Obi-Wan Kenobi
Obi-Wan Kenobi says, "These aren't the droids you're looking for."
*/
package main

import "fmt"

func main() {
	//read input
	fmt.Printf("What is the quote? ")
	var q string
	fmt.Scanln(&q)
	fmt.Printf("Who said it? ")
	var p string
	fmt.Scanln(&p)

	o := p + " says, \"" + q + "\""

	//printing program output
	fmt.Println(o)
}
