/*What is your age? 15
You are not old enough to legally drive.
Or
What is your age? 35
You are old enough to legally drive*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {
	p := prompt.New()

	a := p.ScanInt("What is your age?")

	if a >= 16 {
		fmt.Println("You are old enough to legally drive.")
	} else {
		fmt.Println("You are not old enough to legally drive.")
	}
}
