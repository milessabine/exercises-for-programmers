/* What's the minimum length? 8
How many special characters? 2
How many numbers? 2
Your password is
aurn2$ls#
*/

package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/milessabine/prompt"
)

func main() {

	var letters = []string{"z", "x", "c", "v", "b", "n", "m", "a", "s", "d", "f", "g", "h", "j", "k", "l", "q", "w",
		"e", "r", "t", "y", "u", "i", "o", "p"}
	var numbers = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9", "0"}
	var special = []string{"!", "@", "#", "$", "%", "^", "&", "*"}
	var password []string

	p := prompt.New()

	l, _ := p.ScanInt("What is the length?")
	c, _ := p.ScanInt("How many special characters?")
	n, _ := p.ScanInt("How many numbers?")

	rand.Seed(time.Now().Unix())
	for i := 0; i < l; i++ {
		randLetter := rand.Intn(len(letters))
		password = append(password, letters[randLetter])
	}
	for i := 0; i < n; i++ {
		randNumber := rand.Intn(len(numbers))
		password = append(password, numbers[randNumber])
	}

	for i := 0; i < c; i++ {
		randCharacter := rand.Intn(len(special))
		password = append(password, special[randCharacter])
	}
	fmt.Println(password)
	for range password {
		i := rand.Intn(len(password))
		j := rand.Intn(len(password))
		password[i], password[j] = password[j], password[i]
	}
	fmt.Println("Your password is:")
	fmt.Println(strings.Join(password, ""))
}
