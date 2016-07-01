/*Enter a noun: dog
enter a verb: walk
enter an adjective: blue
enter an adverb: quickly
Do you walk your blue dog quickly? That's hilarious!*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	p := prompt.New()

	n := p.Scan("Enter a noun")
	v := p.Scan("Enter a verb")
	aj := p.Scan("Enter an adjective")
	av := p.Scan("Enter an adverb")

	printMadLib(n, v, aj, av)
}

/*func promptAndScan(n string, scanner *bufio.Scanner) string {
	fmt.Print(n, ": ")
	scanner.Scan()
	return scanner.Text()
}*/

func printMadLib(noun string, verb string, adjective string, adverb string) {
	fmt.Printf("Do you %s your %s %s %s? That's hilarious!\n", verb, adjective, noun, adverb)
}
