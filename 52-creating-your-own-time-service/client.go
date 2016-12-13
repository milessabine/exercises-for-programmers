//localhost:8080

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Time struct {
	CurrentTime time.Time `json:"currentTime"`
}

func main() {
	resp, err := http.Get("http://localhost:8080/")
	if err != nil {
		log.Fatal("Unable to reach data.")
	}
	defer resp.Body.Close()

	dec := json.NewDecoder(resp.Body)
	var t Time
	err = dec.Decode(&t)
	if err != nil {
		log.Fatal("Unable to decode data.")
	}
	//Mon Jan 2 15:04:05 -0700 MST 2006
	// 15:06:26 UTC January 4 2050.
	fmt.Println(t.CurrentTime.UTC().Format("15:04:05 MST January 2 2006"))

}
