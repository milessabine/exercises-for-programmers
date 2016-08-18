/*What is your balance? 5000
What is the APR on the card (as a percent)? 12
What is the monthly payment you can make? 100
It will take you 70 months to pay off this card.
n= number of months
i= daily rate (APR/365)
b= balance
p= monthly payment
*/

package main

import (
	"fmt"
	"math"

	"github.com/milessabine/prompt"
)

func main() {
	p := prompt.New()

	b := p.ScanFloat64("What is your balance?")
	apr := p.ScanFloat64("What is the APR on the card (as a percent)?")
	m := p.ScanFloat64("What is the monthly payment you can make?")

	n := payOff(apr, b, m)

	fmt.Printf("It will take you %d months to pay off this card.\n", n)

}

func payOff(apr float64, b float64, p float64) int {
	i := (apr / 100) / 365
	z := math.Pow(1+i, 30)
	x := math.Log(1 + (b/p)*(1-z))
	y := math.Log(1 + i)

	n := (-1.0 / 30) * (x / y)

	return int(math.Ceil(n))
}
