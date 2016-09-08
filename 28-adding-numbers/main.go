/* Prompt for 5 numbers and Print the sum.
Enter a number :1
Enter a number: 2
Enter a number: 3
Enter a number: 4
Enter a numbr: 5
The total is 15
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {
	p := prompt.New()
	a := 0
	sum := 0
	for i := 1; i <= 5; i++ {
		a = p.ScanInt("Enter a number:")
		sum = a + sum
	}

	fmt.Printf("The total is %d.\n", sum)

}
