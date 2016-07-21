/*What is the order amount? 10
What state do you live in? Wisconsin
The tax is $0.50
The total is $10.50
Wisconsin tax= %5
For Wisconsin reseidents prompt for county.
Eau claire= 0.005 + taxrate
Dunn= 0.004+ taxrate
Illinois= %8
all other states = no tax*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {
	p := prompt.New()

	p1 := p.ScanFloat64("What is the order amount?")
	loc := p.Scan("What state do you live in?")

	taxrate := 0.0
	if loc == "Wisconsin" {
		taxrate = 0.05
		c := p.Scan("What is your county of residence?")
		if c == "Eau Claire" {
			taxrate = 0.055
		} else if c == "Dunn" {
			taxrate = 0.054
		}
	} else if loc == "Illinois" {
		taxrate = 0.08
	}

	tax := p1 * taxrate
	tot := p1 + tax

	fmt.Printf("The tax is $%0.2f\n", tax)
	fmt.Printf("The total is $%0.2f\n", tot)

}
