/* Enter a name: Homer/ Bart/ Maggie/ Lisa/ Moe
The winner is... Maggie
Use a loop to capture input into an array.
Exit loop when user leaves blank entry.
*/

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/milessabine/prompt"
)

var names []string

func main() {
	p := prompt.New()

	for i := 0; ; i++ {
		n := p.Scan("Enter a name:")
		if n == "" {
			break
		} else {
			names = append(names, n)
		}

	}

	rand.Seed(time.Now().Unix())
	number := rand.Intn(len(names))
	fmt.Printf("The winner is ... %s", names[number])
}
