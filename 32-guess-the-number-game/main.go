/* Let's play Guess the Number.
Pick a difficulty level (1, 2, or 3) : 1
I have my number. What's your guess? 1
Too low. Guess again: 5
Too high. Guess again: 2
You got it in 3 guess!
Play again? n
Goodbye!
Difficulty 1 = 1-10 | 2 = 1-100 | 3 = 1-1000

*/

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/milessabine/prompt"
)

func main() {

	p := prompt.New()

	fmt.Println("Let's play Guess the Number.")

	for {
		play(p)
		x := p.Scan("Play again?")
		if x == "n" {
			break
		}
	}
	fmt.Println("Goodbye.")
}

func play(p prompt.Prompt) {
	d, err := p.ScanInt("Pick a difficulty level (1, 2, or 3): ")
	if err != nil {
		fmt.Println("Invalid difficulty level. Please try again")
		return
	}
	var n int
	if d == 1 {
		n = 10
	} else if d == 2 {
		n = 100
	} else if d == 3 {
		n = 1000
	}
	rand.Seed(time.Now().Unix())
	n = rand.Intn(n) + 1

	g, err := p.ScanInt("I have my number. What's your guess?: ")

	for i := 1; ; i++ {
		if err != nil {
			g, err = p.ScanInt("Not a valid number. Guess again:")
			continue
		}
		if g < n {
			g, err = p.ScanInt("Too low. Guess again:")

		} else if g > n {
			g, err = p.ScanInt("Too high. Guess again:")

		} else {
			fmt.Printf("You got it in %d guesses!\n", i)
			break
		}

	}

}
