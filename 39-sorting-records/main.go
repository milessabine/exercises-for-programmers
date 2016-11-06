/* Create a program that sorts employees by last name
Name             | Position | Separation Date
Jacquelyn Jackson|	DBA     |
*/

package main

import (
	"fmt"
	"sort"
)

type Employee struct {
	FirstName      string
	LastName       string
	Position       string
	SeparationDate string
}
type Employees []Employee

func (e Employee) fullName() string {
	return e.FirstName + " " + e.LastName
}

// Len is the number of elements in the collection.
func (e Employees) Len() int {
	return len(e)
}

// Less reports whether the element with
// index i should sort before the element with index j.
func (e Employees) Less(i, j int) bool {
	if e[i].LastName < e[j].LastName {
		return true
	}
	return false
}

// Swap swaps the elements with indexes i and j.
func (e Employees) Swap(i, j int) {
	e[i], e[j] = e[j], e[i]
}

func main() {
	var employees Employees

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
	fmt.Println(employees)
	sort.Sort(employees)
	fmt.Println(employees)
}
