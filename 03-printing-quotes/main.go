/*
What is the quote? These arn't the droids you're looking for.
Who said it? Obi-Wan Kenobi
Obi-Wan Kenobi says, "These aren't the droids you're looking for."
*/
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	//read input
	fmt.Printf("What is the quote? ")
	scanner.Scan()
	q := scanner.Text()
	fmt.Printf("Who said it? ")
	scanner.Scan()
	p := scanner.Text()
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}

	o := p + " says, \"" + q + "\""

	//printing program output
	fmt.Println(o)
}
