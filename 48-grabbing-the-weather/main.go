/* Use OpenWeatherMap http://openweathermap.org/current
Output:
Where are you? Chicago IL
Chicago weather:
65 degrees Fahrnheit
*/

package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/milessabine/prompt"
)

type Weather struct {
	X Main `json:"main"`
}
type Main struct {
	Temp float64 `json:"temp"`
}

func main() {
	id := flag.String("appid", "", "app identifier")
	flag.Parse()
	p := prompt.New()
	city := p.Scan("Where are you?")
	url := "http://api.openweathermap.org/data/2.5/weather?q=" + city + "&units=imperial&appid=" + *id

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal("Unable to find address.")
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var b Weather
	err = dec.Decode(&b)
	if err != nil {
		log.Fatal("Unable to decode.", err)
	}

	fmt.Printf("%s Weather:\n", strings.Title(city))
	fmt.Printf("%0.1f degrees Farhenheit\n", b.X.Temp)
}
