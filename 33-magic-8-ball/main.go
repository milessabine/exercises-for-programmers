/* What's our question? Will I be rich and famous?
Ask again later.
Potential Answers:
"Yes"
"No"
"Maybe"
"Ask again later"
*/

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/milessabine/prompt"
)

var answers = []string{"Yes", "No", "Maybe", "Ask again later."}

func main() {
	p := prompt.New()

	p.Scan("What's your question?")

	rand.Seed(time.Now().Unix())
	n := rand.Intn(4)
	fmt.Println(answers[n])
}
