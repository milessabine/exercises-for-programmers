/*Given an input file, read the file and look for the word "utilize."
Replace with "use."
Input: One should never utilize the word "utilize" in writing. Use "use" instead.
Output: One should never use the word "use" in writing. Use "use" instead.
*/

package main

import (
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func main() {

	f, err := os.Open("data.txt")
	if err != nil {
		log.Fatal("Unable to open file")
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("Unable to read file.", err)
	}

	s := strings.Replace(string(b), "utilize", "use", -1)

	of, err := os.Create("newdata.txt")
	if err != nil {
		log.Fatal("Unable to create file.", err)
	}
	defer of.Close()

	_, err = of.WriteString(s)
	if err != nil {
		log.Fatal("Unable to write file.")
	}

}
