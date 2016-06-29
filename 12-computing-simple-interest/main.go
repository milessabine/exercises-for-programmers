/*Enter the principal: 1500
Enter the rate of interest: 4.3
Enter the number of years: 4
After 4 years t 4.3%, the investment will be worth $1758
interest formula A=P(1+rt)
P= principal amount
r= annual rate of interest
t= number of years invested
A= amount at end of investment */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	p := promptAndScanInt("Enter the principal:", scanner)
	r := promptAndScanFloat64("Enter the rate of interest:", scanner)
	t := promptAndScanInt("Enter the number of years:", scanner)
	a := float64(p) * (1 + (r/100)*float64(t))

	fmt.Printf("After %d years at %0.1f%%, the investment will be worth $%0.2f\n", t, r, a)

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
		fmt.Println("Not a valid number", s, err)
		return 0
	}
	return i
}

func promptAndScanFloat64(n string, scanner *bufio.Scanner) float64 {
	s := promptAndScan(n, scanner)
	f, err := strconv.ParseFloat(s, 64)
	if err != nil {
		fmt.Println("Please input valid number", s, err)
		return 0.0
	}
	return f
}
