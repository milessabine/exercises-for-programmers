/* What is the length of the room in feet? 15
What is the width of the room in feet? 20
You entered dimensions of 15 by 20 feet.
The area is
300 square feet
27.871 square meters */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const conversion = 0.09290304

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	l := promptAndScanInt("What is the length of the room in feet?", scanner)
	w := promptAndScanInt("What is the width of the room in feet?", scanner)
	sqf := l * w
	sqm := float64(sqf) * conversion

	fmt.Printf("You entered the dimensions of %d by %d feet.\n", l, w)
	fmt.Println("The area is:")
	fmt.Printf("%d square feet \n", sqf)
	fmt.Printf("%0.3f square meters \n", sqm)

}

func promptAndScan(n string, scanner *bufio.Scanner) string {
	fmt.Print(n, " ")
	scanner.Scan()
	return scanner.Text()
}

func promptAndScanInt(n string, scanner *bufio.Scanner) int {
	s := promptAndScan(n, scanner)
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Cannot use this number", s, err)
		return 0
	}
	return i
}
