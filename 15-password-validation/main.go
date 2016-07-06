/*What is the password? 12345
I don't know you.
Or
What is the password? abc$123
Welcome!*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

const password = "abc$123"

func main() {
	p := prompt.New()

	c := p.Scan("What is the password?")

	if checkPassword(c) {
		fmt.Println("Welcome!")
	} else {
		fmt.Println("I don't know you.")
	}

}

func checkPassword(p string) bool {
	if p == password {
		return true
	}
	return false

}
