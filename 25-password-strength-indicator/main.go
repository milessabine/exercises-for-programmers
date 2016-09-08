/* Password strength
very weak = only numbers, <8 characters
weak = only letters, <8 characters
strong = letters, at least one number, >= 8 characters
very strong = letters, numbers, special characters, >= 8 characters
Output
The password '12345' is a very weak password.
The password 'abcdef' is a weak password.
The password 'abc123xyz' is a strong password.
The password '1337h@xor!' is a very strong password.
*/

package main

import (
	"fmt"
	"unicode"

	"github.com/milessabine/prompt"
)

type passwordStrength int

const (
	veryWeak passwordStrength = iota
	weak
	strong
	veryStrong
)

func main() {
	p := prompt.New()

	a := p.Scan("Enter password:")

	s := passwordValidator(a)
	var str string
	switch s {
	case veryWeak:
		str = "very weak"
	case weak:
		str = "weak"
	case strong:
		str = "strong"
	case veryStrong:
		str = "very strong"
	}
	fmt.Printf("The password '%s' is a %s password.\n", a, str)
}

func passwordValidator(pass string) passwordStrength {

	l := len(pass)
	if l < 8 {
		num := true
		for _, r := range pass {
			d := unicode.IsNumber(r)
			if d == false {
				num = false
				break
			}
		}
		if num == true {
			return veryWeak
		}

		let := true
		for _, s := range pass {
			f := unicode.IsLetter(s)
			if f == false {
				let = false
				break
			}
		}
		if let == true {
			return weak
		}
	}

	special := false
	let := false
	num := false

	for _, r := range pass {
		if unicode.IsNumber(r) {
			num = true
		} else if unicode.IsLetter(r) {
			let = true
		} else if !unicode.IsNumber(r) && !unicode.IsLetter(r) {
			special = true
		}
	}
	if !let || !num {
		return weak
	}
	if special == true {
		return veryStrong
	}
	return strong

}
