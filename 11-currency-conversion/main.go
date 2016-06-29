/*How many euros are you exchanging? 81
What is the exchange rate? 137.51
81 euros at an exchange rat of 137.51 is
111.38 U.S. dollars.
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	eur := promptAndScanInt("How many euros are you exchanging?", scanner)
	ex := promptAndScanFloat64("What is the exchange rate?", scanner)
	us := float64(eur) * ex
	//fmt.Printf("%d euros at an exhange rate of %0.2f is:\n", eur, ex)
	//fmt.Printf("%0.2f U.S. dollars.\n", float64(us)/100)
	//Single output statement???
	fmt.Printf("%d euros at an exchange rate of %0.2f is:\n%0.2f U.S. dollars.\n",
		eur, ex, float64(us))
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
