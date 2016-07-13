/* Prompt for weight, gender, number of drinks,amount of alcohol, amount of time since last drink.
BAC=(A x 5.14 / W x r) - 0.015 x H
A= total alcohol consumed in oz.
W= body weight
r= alcohol distribution ratio: 0.73 for men, 0.66 for women
H= number of hours since last drink.
Your BAC is 0.08
It is not legal for you to drive.
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

const male = 0.73
const female = 0.66

func main() {
	p := prompt.New()

	a := p.ScanInt("How many ounces of alcohol did you consume?")
	w := p.ScanInt("How many pounds do you weigh?")
	g := p.Scan("What is your gender?")
	h := p.ScanInt("How many hours has it been since your last drink?")
	var r float64

	if g == "female" {
		r = female
	} else if g == "male" {
		r = male
	} else {
		fmt.Println("Please enter male or female.")
		return
	}

	bac := ((float64(a) * 5.14 / float64(w) * r) - 0.015*float64(h))
	fmt.Printf("Your BAC is %0.2f\n", bac)
	if bac >= 0.08 {
		fmt.Println("It is not legal for you to drive.")
	} else {
		fmt.Println("It is legal for you to drive.")
	}
}

/*if g == "female" {
  r := 0.73
}
if g == "male" {
  r := 0.66
}*/
