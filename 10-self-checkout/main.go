/*Enter the price of item 1: 25
Enter the quantity of item 1:2
Enter the price of item 2 :10
Enter the quantity of item 2: 1
Enter the price of item 3: 4
Enter the quantity of item 3:1
Subtotal:$64.00
Tax: 3.52
Total: $67.52
*/

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const taxRate = 0.055

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	p1 := promptAndScanInt("Enter the price of item 1:", scanner)
	q1 := promptAndScanInt("Enter the quantity of item 1:", scanner)
	p2 := promptAndScanInt("Enter the price of item 2:", scanner)
	q2 := promptAndScanInt("Enter the quantity of item 2:", scanner)
	p3 := promptAndScanInt("Enter the price of item 3:", scanner)
	q3 := promptAndScanInt("Enter the quantity of item 3:", scanner)

	t1 := q1 * p1 * 100
	t2 := q2 * p2 * 100
	t3 := q3 * p3 * 100
	sub := t1 + t2 + t3
	tax := int(float64(sub) * taxRate)
	tot := tax + sub

	fmt.Printf("Subtotal: $%0.2f\n", float64(sub)/100)
	fmt.Printf("Tax: $%0.2f\n", float64(tax)/100)
	fmt.Printf("Total: $%0.2f\n", float64(tot)/100)

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
		fmt.Println("Please input valid number", s, err)
		return 0
	}
	return i
}
