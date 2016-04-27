/*what is your name? Brian
Hello, Brian,nice to meet you!*/
package main

import "fmt"

func main() {
	n := getInput()
	o := "hello, " + n + ", nice to meet you!"
	printOutput(o)
}
func printOutput(s string) { fmt.Println(s) }
func getInput() string {
	fmt.Printf("What is your name? ")
	str := ""
	if _, err := fmt.Scan(&str); err != nil {
		fmt.Println(err)
	}
	return str
}
