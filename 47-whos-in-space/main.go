/* Read data from API in a tabular format.
http://api.open-notify.org/astros.json
Output:
There are x people in space right now:
Name | Craft
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name  string
	Craft string
}

func (p Person) String() string {
	return fmt.Sprintf(p.Name + "     | " + p.Craft)
}

type Base struct {
	People []Person
}

func main() {

	//	f, err := os.Open("data.json")
	//	if err != nil {
	//		log.Fatal("Unable to open file", err)
	//}
	//	defer f.Close()

	resp, err := http.Get("http://api.open-notify.org/astros.json")
	if err != nil {
		log.Fatal("Unable to find address.")
	}
	defer resp.Body.Close()

	// Need to convert body type []byte to io.Reader
	dec := json.NewDecoder(resp.Body)
	var b Base
	err = dec.Decode(&b)
	if err != nil {
		log.Fatal("Unable to decode.", err)
	}
	fmt.Printf("There are %d people in space right now:\n", len(b.People))
	fmt.Println("Name                  | Craft")
	fmt.Println("----------------------|-------")
	for _, people := range b.People {
		fmt.Println(people)
	}
}
