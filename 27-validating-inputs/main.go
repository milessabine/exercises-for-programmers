/*
Enter the first name: J
Enter the last name:
Enter the ZIP code: ABCDE
Enter an employee ID: A12-1234
"J" is not a valid first name. It is too short.
The last name must be filled in.
The ZIP code must be numeric.
A12-1234 is not a valid ID.

OR

Enter the first name: Jimmy
Enter the last name: James
Enter the ZIP code: 55555
Enter an employee ID: TK-421
There were no errors found.

Rules:
- The first name must be filled in
- THe last name must be filled in
- The first and last names must be at least two characters long
- An employee ID is in the format AA-1234. So, two letters, a hyphen, and four numbers.
- The ZIP code must be a number.
*/
package main

import (
	"fmt"
	"regexp"

	"github.com/milessabine/prompt"
)

var idregex = regexp.MustCompile(`\A[A-Za-z]{2}-[0-9]{4}\z`)

func main() {
	p := prompt.New()

	fn := p.Scan("Enter the first name:")
	ln := p.Scan("Enter the last name:")
	_, zcerr := p.ScanInt("Enter the ZIP code:")
	id := p.Scan("Enter an employee ID:")

	var errors bool

	if fn == "" {
		fmt.Println("The first name must be filled in.")
		errors = true
	} else if len(fn) < 2 {
		fmt.Printf("\"%s\" is not a valid first name. It is too short.\n", fn)
		errors = true
	}

	if ln == "" {
		fmt.Println("The last name must be filled in.")
		errors = true
	} else if len(ln) < 2 {
		fmt.Printf("\"%s\" is not a valid last name. It is too short.\n", ln)
		errors = true
	}

	if zcerr != nil {
		fmt.Println("The ZIP code must be numeric.")
		errors = true
	}

	if !idregex.MatchString(id) {
		fmt.Printf("\"%s\" is not a valid ID.\n", id)
		errors = true
	}

	if !errors {
		fmt.Println("There were no errors found")
	}
}
