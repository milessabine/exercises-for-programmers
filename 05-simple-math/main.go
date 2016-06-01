/*What is the first number? 10
What is the second number? 5
10+5= 15
10-5+=5
10*5= 50
10/5= 2 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	f := promptAndScan("What is the first number?", scanner)
	fn, err := strconv.Atoi(f)
	if err != nil {
		fmt.Println("Cannot convert first number", f, err)
		return
	}

	s := promptAndScan("What is the second number?", scanner)
	sn, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println("Cannot convert second number", s, err)
		return
	}

	fmt.Printf("%d + %d = %d \n", fn, sn, fn+sn)
	fmt.Printf("%d - %d = %d \n", fn, sn, fn-sn)
	fmt.Printf("%d * %d = %d \n", fn, sn, fn*sn)
	fmt.Printf("%d / %d = %0.f \n", fn, sn, float64(fn)/float64(sn))
}

func promptAndScan(n string, scanner *bufio.Scanner) string {
	fmt.Print(n, " ")
	scanner.Scan()
	return scanner.Text()
}
