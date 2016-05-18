/*Enter a noun: dog
enter a verb: walk
enter an adjective: blue
enter an adverb: quickly
Do you walk your blue dog quickly? That's hilarious!*/

package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Print("Enter a noun: ")
	scanner.Scan()
	n := scanner.Text()
	fmt.Print("Enter a verb: ")
	scanner.Scan()
	v := scanner.Text()
	fmt.Print("Enter an adjective: ")
	scanner.Scan()
	aj := scanner.Text()
	fmt.Print("Enter an adverb: ")
	scanner.Scan()
	av := scanner.Text()

	printMadLib(n, v, aj, av)
}

func printMadLib(noun string, verb string, adjective string, adverb string) {
	fmt.Printf("Do you %s your %s %s %s? That's hilarious!\n", verb, adjective, noun, adverb)
}
