/* Create multiplication table for numbers 0 through 12
Output:
0 x 0 = 0
0 x 1 = 0
*/

package main

import "fmt"

func main() {

	/*	for x, y := 0, 0; x <= 12 && y <= 12; y++ {

		if y == 12 {
			x++
			y = 0
		}
		fmt.Printf("%d x %d = %d\n", x, y, x*y)
	} */

	fmt.Println(" --- --- --- --- --- --- --- --- --- --- --- --- --- --- ")
	fmt.Println("|   | 0 | 1 | 2 | 3 | 4 | 5 | 6 | 7 | 8 | 9 | 10| 11| 12|")

	for y := 0; y <= 12; y++ {
		fmt.Println(" --- --- --- --- --- --- --- --- --- --- --- --- --- --- ")
		fmt.Printf("|%3d|", y)
		for x := 0; x <= 12; x++ {
			fmt.Printf("%3d|", x*y)
		}
		fmt.Println("")
	}
	fmt.Println(" --- --- --- --- --- --- --- --- --- --- --- --- --- --- ")
}
