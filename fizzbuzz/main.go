/*Count up from 1 to 100
For every number that divides evenly by three, print fizz instead of number.
For every number that divides evenly by five, print buzz instead of number
For every number that divides evenly by 3 and 5, print fizzbuzz instead of number.
*/

package main

import "fmt"

func main() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 && i%5 == 0 {
			fmt.Println("fizzbuzz")
		} else if i%5 == 0 {
			fmt.Println("buzz")
		} else if i%3 == 0 {
			fmt.Println("fizz")
		} else {
			fmt.Println(i)
		}

	}
}
