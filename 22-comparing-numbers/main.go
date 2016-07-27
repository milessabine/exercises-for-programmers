/*Enter the first number: 1
Enter the second number: 51
Enter the third number: 2
The Largest number is 51
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {
	p := prompt.New()

	n1 := p.ScanInt("Enter the first number:")
	n2 := p.ScanInt("Enter the second number:")
	n3 := p.ScanInt("Enter the third number:")
	tot := 0
	if n1 > n2 && n1 > n3 {
		tot = n1
	} else if n2 > n1 && n2 > n3 {
		tot = n2
	} else if n3 > n1 && n3 > n2 {
		tot = n3
	}
	fmt.Printf("The largest number is %d\n", tot)
}
