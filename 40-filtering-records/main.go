/*Enter a search string: Jac
Results:
Name|Position|Separation Date
Jacquelyn Jackson...
Jake Jacobson
*/

package main

import (
	"fmt"
	"strings"

	"github.com/milessabine/prompt"
)

type Employee struct {
	FirstName      string
	LastName       string
	Position       string
	SeparationDate string
}

func main() {
	var employees []Employee

	john := Employee{
		FirstName:      "John",
		LastName:       "Johnson",
		Position:       "Manager",
		SeparationDate: "2016-12-31",
	}
	tou := Employee{
		FirstName:      "Tou",
		LastName:       "Xiong",
		Position:       "Software Engineer",
		SeparationDate: "2016-10-05",
	}
	michaela := Employee{
		FirstName:      "Michaela",
		LastName:       "Michaelson",
		Position:       "District Manager",
		SeparationDate: "2015-12-19",
	}
	jake := Employee{
		FirstName:      "Jake",
		LastName:       "Jacobson",
		Position:       "Programmer",
		SeparationDate: "",
	}
	jacquelyn := Employee{
		FirstName:      "Jacquelyn",
		LastName:       "Jackson",
		Position:       "DBA",
		SeparationDate: "",
	}
	sally := Employee{
		FirstName:      "Sally",
		LastName:       "Weber",
		Position:       "Web Developer",
		SeparationDate: "",
	}

	employees = append(employees, john, tou, michaela, jake, jacquelyn, sally)

	p := prompt.New()
	s := p.Scan("Enter a search string:")

	for _, x := range employees {
		if strings.Contains(x.FirstName, s) || strings.Contains(x.LastName, s) {
			fmt.Println(x.fullName())
		}
	}

}

func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}
