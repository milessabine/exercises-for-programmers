/*Total of 7 Names
Ling, Mai
Johnson, Jim
Jones, Aaron
Jones, Chris
Swift, Geoffrey
Xiong, Fong
Zarnecki, Sabrina
*/

package main

import (
	"fmt"
	"log"
	"os"
	"sort"
)

type Name struct {
	FirstName string
	LastName  string
}
type Names []Name

func (n Name) fullName() string {
	return n.LastName + ", " + n.FirstName
}

func (n Names) Len() int {
	return len(n)
}

func (n Names) Less(i, j int) bool {
	if n[i].LastName < n[j].LastName {
		return true
	}
	return false
}

func (n Names) Swap(i, j int) {
	n[i], n[j] = n[j], n[i]
}

func main() {
	var names Names

	mai := Name{
		FirstName: "Mai",
		LastName:  "Ling",
	}
	jim := Name{
		FirstName: "Jim",
		LastName:  "Johnson",
	}
	aaron := Name{
		FirstName: "Aaron",
		LastName:  "Jones",
	}
	chris := Name{
		FirstName: "Chris",
		LastName:  "Jones",
	}
	geoffrey := Name{
		FirstName: "Geoffrey",
		LastName:  "Swift",
	}
	fong := Name{
		FirstName: "Fong",
		LastName:  "Xiong",
	}
	sabrina := Name{
		FirstName: "Sabrina",
		LastName:  "Zarnecki",
	}
	names = append(names, mai, jim, aaron, chris, geoffrey, fong, sabrina)

	sort.Sort(names)
	f, err := os.Create("names.txt")
	if err != nil {
		log.Fatal("Unable to create file.", err)
	}
	defer f.Close()

	fmt.Fprintf(f, "Total of %d names\n", len(names))

	for _, name := range names {
		if _, err := f.WriteString(name.fullName() + "\n"); err != nil {
			log.Fatal("Unable to write string to the file.", err)
		}
	}

}
