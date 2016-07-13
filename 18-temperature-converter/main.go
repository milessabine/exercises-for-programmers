/* Press C to convert from Farenheit to Celius.
Press F to convert from Celcius to Farnheit.
Your choice: C
Please enter the temperature in Farenheit: 32
The temperature in Celsiusis: 0
C= (F-32) x 5/9
F= (C x 9/5) + 32*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {
	p := prompt.New()

	fmt.Println("Press C to convert from Farenheit to Celcius.")
	fmt.Println("Press F to convert from Celcius to Farenheit.")

	d := p.Scan("Choose C or F.")
	var in string
	var out string
	var fn func(float64) float64

	if d == "C" || d == "c" {
		in = "Farenheit"
		out = "Celcius"
		fn = fToC
	} else if d == "F" || d == "f" {
		in = "Celcius"
		out = "Farenheit"
		fn = cToF
	} else {
		fmt.Println("Please enter C or F.")
		return
	}
	t := p.ScanFloat64("Please enter the temperature in " + in + ":")
	t = fn(t)
	fmt.Printf("The temperature in %s is: %f\n", out, t)
}

func cToF(c float64) float64 {
	return (c * 9 / 5) + 32
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}
