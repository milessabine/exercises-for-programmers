/*How many people? 8
How many pizzas do you have? 2
8 People with 2 pizzas
Each person gets 2 slices of pizza.
There are 0 leftover pieces.
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

	p := promptAndScanInt("How many people?", scanner)
	pie := promptAndScanInt("How many pizzas do you have?", scanner)
	s := promptAndScanInt("How many slices per pizza?", scanner)
	tot := pie * s
	spp := tot / p
	l := tot % p

	fmt.Printf("%d people with %d pizzas\n", p, pie)
	fmt.Printf("Each person gets %d pieces of pizza.\n", spp)
	fmt.Printf("There are %d leftover pieces.\n", l)

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
