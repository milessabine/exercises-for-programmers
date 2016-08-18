/*Enter two strings and I'll tell you if they are anagrams:
Enter the first string: note
Enter the second string: tone
"note" and "tone" are anagrams.
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {
	p := prompt.New()

	fmt.Println("Enter the first two strings and I'll tell you if they are anagrams")
	a := p.Scan("Enter the first string:")
	b := p.Scan("Enter the second string")

	c := isAnagram(a, b)
	if c == true {
		fmt.Printf("\"%s\" and \"%s\" are anagrams.\n", a, b)
	} else {
		fmt.Printf("\"%s\" and \"%s\" are not anagrams.\n", a, b)

	}

}

func isAnagram(a string, b string) bool {
	if len(a) != len(b) {
		return false
	}

	for _, x := range a {
		found := false

		for _, y := range b {
			if y == x {
				found = true
				break
			}
		}

		if found == false {
			return false
		}
	}

	return true
}
