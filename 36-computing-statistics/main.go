/* Enter a number: 100,200, 1000, 300, done
Numbers: 100, 200, 1000, 300
The average is 400.
The minimum is 100.
The maximum is 1000.
The standard deviation is 400.25

Stop prompting for numbers when user enters "done"
*/

package main

import (
	"fmt"
	"math"
	"strconv"

	"github.com/milessabine/prompt"
)

func main() {

	p := prompt.New()

	var numbers []int

	for {
		number := p.Scan("Enter a number: ")
		if number == "done" {
			break
		}
		n, err := strconv.Atoi(number)
		if err != nil {
			fmt.Println("You entered bad data.")
			continue
		}
		numbers = append(numbers, n)
	}

	sum := 0
	for _, a := range numbers {
		sum += a
	}
	mean := sum / len(numbers)
	fmt.Printf("The average is: %d.\n", mean)

	max := 0
	for _, a := range numbers {
		if a > max {
			max = a
		}
	}
	min := max
	for _, a := range numbers {
		if min > a {
			min = a
		}
	}
	fmt.Printf("The minimum is %d.\n", min)
	fmt.Printf("The maximum is %d.\n", max)

	diff := 0.0
	for _, a := range numbers {
		f := float64(mean - a)
		diff += f * f
	}
	diffMean := diff / float64(len(numbers))
	stndev := math.Sqrt(diffMean)
	fmt.Printf("The standard deviation is %0.2f.\n", stndev)

}
