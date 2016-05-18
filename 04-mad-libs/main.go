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

	n := promptAndScan("Enter a noun", scanner)
	v := promptAndScan("Enter a verb", scanner)
	aj := promptAndScan("Enter an adjective", scanner)
	av := promptAndScan("Enter an adverb", scanner)

	printMadLib(n, v, aj, av)
}

func promptAndScan(n string, scanner *bufio.Scanner) string {
	fmt.Print(n, ": ")
	scanner.Scan()
	return scanner.Text()
}

func printMadLib(noun string, verb string, adjective string, adverb string) {
	fmt.Printf("Do you %s your %s %s %s? That's hilarious!\n", verb, adjective, noun, adverb)
}
