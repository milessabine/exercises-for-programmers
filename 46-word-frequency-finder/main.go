/* Given the text file words.txt with this content:
badger badger badger badger mushroom mushroom snake badger badger badger
Count the occurance of the words:
badger: *******
mushroom: **
snake: *
*/

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"sort"
	"strings"
)

type Word struct {
	word  string
	count int
}

func (w Word) String() string {
	str := w.word + ": "
	for i := 0; i < w.count; i++ {
		str += "*"
	}
	return str
}

type Words []Word

func (w Words) Len() int           { return len(w) }
func (w Words) Swap(i, j int)      { w[i], w[j] = w[j], w[i] }
func (w Words) Less(i, j int) bool { return w[i].count < w[j].count }

func main() {

	f, err := os.Open("words.txt")
	if err != nil {
		log.Fatal("Unable to open file.")
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal("Unable to read file.", err)
	}

	c := string(b)
	badger := Word{
		word:  "badger",
		count: strings.Count(c, "badger"),
	}

	mushroom := Word{
		word:  "mushroom",
		count: strings.Count(c, "mushroom"),
	}

	snake := Word{
		word:  "snake",
		count: strings.Count(c, "snake"),
	}
	words := Words{badger, mushroom, snake}
	sort.Sort(sort.Reverse(words))

	for _, a := range words {
		fmt.Println(a)
	}
}
