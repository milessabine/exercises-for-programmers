/* There are 5 employees:
John Smith
Jackie Jackson
Chris Jones
Amnda Cullen
Jeremy Goodwin

Enter an employee name to remove: Chris Jones

There are 4 employees:
John Smith
Jackie Jackson
Amanda Cullen
Jeremy Goodwin
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

var names = []string{"John Smith", "Jackie Jackson", "Chris Jones",
	"Amanda Cullen", "Jeremy Goodwin"}

func main() {

	initEmployees()

	p := prompt.New()

outerLoop:
	for {
		remove := p.Scan("Enter an employee name to remove: ")

		l := len(names)
		for i, b := range names {
			if b == remove {
				names = append(names[:i], names[i+1:]...)
				initEmployees()
				break outerLoop
			}
		}
		if l == len(names) {
			fmt.Println("Name not present.")
		}

	}
}

func initEmployees() {
	tot := len(names)
	fmt.Printf("There are %d  employees:\n", tot)
	for _, a := range names {
		fmt.Println(a)
	}
	fmt.Println()
}
