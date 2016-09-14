/*What is the rate of return? 0
Sorry. That's not a valid input.
ABC = not valid
4 = It will take 18 years to double your initial investment

years = 72/r
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {

	p := prompt.New()
	var r int
	var err error
	for {
		r, err = p.ScanInt("What is the rate of return?")
		if r != 0 && err == nil {
			break
		}

		fmt.Println("Sorry. That's not a valid input.")
	}
	y := 72 / r
	fmt.Printf("It will take %d years to double your initial investment\n", y)
}
