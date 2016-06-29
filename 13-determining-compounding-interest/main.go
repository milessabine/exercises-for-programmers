/*What is the principal ammount? 1500
What is the rate?4.3
What is the number of years? 6
What is the number of times the interest is compounded per year? 4
$1500 invested at 4.3% for 6 years compounded 4 times per year is $1938.84
A=P(1 + r/n)^nt
*/

package main

import (
	"fmt"
	"math"

	"github.com/milessabine/prompt"
)

func main() {
	p := prompt.New()

	s := p.ScanInt("Enter the principal:")
	r := p.ScanFloat64("Enter the rate of interest:")
	t := p.ScanInt("Enter the number of years:")
	n := p.ScanInt("What is the number of ties the interest is compounded per year?")
	a := float64(s) * math.Pow(1+(r/100)/float64(n), float64(n*t))

	fmt.Printf("$%d invested at %0.1f%% for %d years compounded %d times per year is $%0.2f\n", s, r, t, n, a)

}
