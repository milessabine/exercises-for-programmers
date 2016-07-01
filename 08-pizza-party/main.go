/*How many people? 8
How many pizzas do you have? 2
8 People with 2 pizzas
Each person gets 2 slices of pizza.
There are 0 leftover pieces.
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	p := prompt.New()
	pe := p.ScanInt("How many people?")
	pie := p.ScanInt("How many pizzas do you have?")
	s := p.ScanInt("How many slices per pizza?")
	tot := pie * s
	spp := tot / pe
	l := tot % pe

	fmt.Printf("%d people with %d pizzas\n", pe, pie)
	fmt.Printf("Each person gets %d pieces of pizza.\n", spp)
	fmt.Printf("There are %d leftover pieces.\n", l)

}

/*func promptAndScan(n string, scanner *bufio.Scanner) string {
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
}*/
