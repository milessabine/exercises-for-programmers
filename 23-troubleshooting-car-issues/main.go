/*Prompt
Is the car silent when you turn the key?
If yes: Are the battery terminals corroded
  yes = Clean terminals and try starting again. No = Replace cables and try again
If no: Does the car make a clicking noise?
  yes= Replace the battery
  No: Does the car crank up but fail to start?
    yes= Check the spark plug connections
    No: Does the engine start nd then die?
      yes = Does your car have fuel injection?
        no = Check to ensure the choke is open and closing
        yes = Get it in for service
*/

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

const yes = "y"

func diagnose() string {
	var output string
	p := prompt.New()

	q := p.Scan("Is the car silent when you turn the key?")

	if q == yes {
		q = p.Scan("Are the battery terminals corroded?")
		if q == yes {
			output = "Clean the terminals and try starting again."
		} else {
			output = "Replace cables and try again."
		}
	} else {
		q = p.Scan("Does the car make a clicking noise?")
		if q == yes {
			output = "Replace the battery"
		} else {
			q = p.Scan("Does the car crank up but fail to start?")
			if q == yes {
				output = "Check the spark plug connections."
			} else {
				q = p.Scan("Does the engine start and then die?")
				if q == yes {
					q = p.Scan("Does your car have fuel injection?")
					if q == yes {
						output = "Get it in for service."
					} else {
						output = "Check to ensure the choke is opening and closing."
					}
				}
			}
		}
	}
	return output
}
func main() {
	fmt.Println(diagnose())
}
