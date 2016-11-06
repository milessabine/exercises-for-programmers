/*


 */
package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"strconv"
	"strings"
)

type Employee struct {
	LastName  string
	FirstName string
	Salary    int
}

func (e Employee) String() string {
	return fmt.Sprintf("%s %s %d", e.LastName, e.FirstName, e.Salary)
}

func main() {

	//Open data.csv file
	f, err := os.Open("data.csv")
	if err != nil {
		log.Fatal("Could not open file.", err)
	}
	defer f.Close()
	//Create a variable to hold a slice of employees
	var employees []Employee
	//We will parse each line of data.csv file and create an employee instance
	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("Unable to read file.", err)
	}

	for _, line := range strings.Split(string(b), "\n") {
		line = strings.TrimSpace(line)
		item := strings.Split(line, ",")
		if len(item) < 3 {
			continue
		}
		i, err := strconv.Atoi(item[2])
		if err != nil {
			fmt.Println(err)
			continue
		}
		//Add that employee instance to the slice of employees
		name := Employee{
			LastName:  item[0],
			FirstName: item[1],
			Salary:    i,
		}
		employees = append(employees, name)
	}

	//Print output
	fmt.Println("Last First Salary")
	fmt.Println("------------------")
	for _, employee := range employees {
		fmt.Println(employee.String())
	}
}
