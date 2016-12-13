/* Creates a web service that returns current time as JSON data
Output: The current time is...
*/

package main

import (
	"encoding/json"
	"log"
	"net/http"
	"time"
)

type Time struct {
	CurrentTime time.Time `json:"currentTime"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		t := Time{CurrentTime: time.Now()}
		enc := json.NewEncoder(w)
		err := enc.Encode(t)
		if err != nil {
			log.Println("Unable to encode data.")
		}
	})
	http.ListenAndServe(":8080", nil)

}
