/*What is your age? 25
At what age would you like to retire? 65
You have 40 years left until you can retire.
It's 2015, so you can retire in 2055*/

package main

import (
	"fmt"
	"time"

	"github.com/milessabine/prompt"
)

func main() {
	//scanner := bufio.NewScanner(os.Stdin)
	p := prompt.New()

	ca := p.ScanInt("What is your current age?")
	ra := p.ScanInt("At what age would you like to retire?")
	cy := time.Now().Year()
	yl := ra - ca
	ry := yl + cy

	fmt.Printf("You have %d years left until you can retire.\n", yl)
	fmt.Printf("It's %d, so you can retire in %d.\n", cy, ry)
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
