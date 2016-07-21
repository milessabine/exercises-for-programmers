/* Please entr the nmber of the month : 3
The name of the month is March.
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {
	p := prompt.New()

	m := p.ScanInt("Please enter the number of the month:")
	n := ""
	switch m {
	case 1:
		n = "January"
	case 2:
		n = "February"
	case 3:
		n = "March"
	case 4:
		n = "April"
	case 5:
		n = "May"
	case 6:
		n = "June"
	case 7:
		n = "July"
	case 8:
		n = "August"
	case 9:
		n = "September"
	case 10:
		n = "October"
	case 11:
		n = "November"
	case 12:
		n = "December"
	default:
		fmt.Println("Number was out of range", m)
		return
	}
	fmt.Printf("The name of the month is %s\n", n)
}
