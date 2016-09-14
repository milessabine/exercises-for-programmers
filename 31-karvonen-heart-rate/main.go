/*Resting pulse: 65 Age:22
Intensity | Rate
55% 138 bPM
60% 145 bpm
65% 151bpm
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {

	p := prompt.New()

	rp, err := p.ScanFloat64("Enter your resting heart rate:")
	if err != nil {
		fmt.Println(err)
		return
	}
	a, err := p.ScanFloat64("Enter your age:")
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Resting Pulse: %0.0f \tAge: %0.0f\n\n", rp, a)

	fmt.Println("Intensity | Rate")
	fmt.Println("----------|--------")
	for i := 55; i <= 95; i += 5 {

		tgh := (((220 - a) - rp) * (float64(i) / 100)) + rp
		fmt.Printf("%d%%       | %0.0f bpm\n", i, tgh)
	}

}
