/*What is the order amount>? 10
What is the state? WI
The subtotal is $10.00
The tax is $0.55.
The total is $10.55
Or
What is the order amount? 10
What is the state? MN
The total is $10.00
WI tax= %5.5 No tax for non residents of WI
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

const taxRate = 0.055

func main() {

	p := prompt.New()

	p1 := p.ScanFloat64("What is the order amount?")
	loc := p.Scan("What is the state?")

	tot := p1

	if loc == "WI" {
		tax := p1 * taxRate
		tot = tot + tax
		fmt.Printf("The subtotal is $%0.2f\n", p1)
		fmt.Printf("The tax is $%0.2f\n", tax)
	}

	fmt.Printf("The total is $%0.2f\n", tot)
}
