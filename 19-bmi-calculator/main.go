/*Prompt for weight and height.
Your BMI is 19.5.
You are within the ideal weight range.
Or
Your BMI is 32.5.
You are overweight. You should see your doctor.
bmi= (weight / (height * height)) *703 */

package main

import (
	"fmt"

	"github.com/milessabine/prompt"
)

func main() {
	p := prompt.New()

	w := p.ScanFloat64("What is your weight in lbs?")
	h := p.ScanFloat64("What is your height inches?")

	bmi := (w / (h * h)) * 703
	fmt.Printf("Your BMI is %0.1f\n", bmi)
	if bmi < 18.5 {
		fmt.Println("You are underweight. You should see your doctor.")
	} else if bmi >= 18.5 && bmi <= 25 {
		fmt.Println("You are within the ideal weight range")
	} else if bmi > 25 {
		fmt.Println("You are overweight. You should see your doctor.")
	}
}
