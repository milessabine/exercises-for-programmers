/* Enter a list of numbers, separated by spaces: 1 2 3-8
The even numbers are 2 4 6 8
*/

package main

import (
	"fmt"
	"strconv"

	"github.com/milessabine/prompt"
)

func main() {
	var allNumbers []int
	p := prompt.New()

	n := p.Scan("Enter a list of numbers, separated by spaces:")

	for _, i := range n {
		x := string(i)
		if x == " " {
			continue
		}
		y, err := strconv.Atoi(x)
		if err != nil {
			fmt.Printf("This character is not a valid number. %s\n", x)
			return
		}
		if y%2 == 0 {
			allNumbers = append(allNumbers, y)
		}
	}
	for _, a := range allNumbers {
		fmt.Printf("%d ", a)
	}
	fmt.Println()
}

//func filterEvenNumbers([]string) []string {

//}
