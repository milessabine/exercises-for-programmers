/*https://api.flickr.com/services/feeds/photos_public.gne?format=json&tags=nature
  Create program with graphical interface that takes a string search and displays
  matching pictures.
Remove : jsonFlickrFeed(
*/

package main

import (
	"encoding/json"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

type Base struct {
	Items []Item `json:"items"`
}
type Item struct {
	Media Media `json:"media"`
}

type Media struct {
	URL string `json:"m"`
}

func main() {

	resp, err := http.Get("https://api.flickr.com/services/feeds/photos_public.gne?format=json&tags=nature")
	if err != nil {
		log.Fatal("Unable to reach URL.")
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal("Unable to read response body.")
	}
	t := strings.TrimPrefix(string(b), "jsonFlickrFeed(")
	t = strings.TrimSuffix(t, ")")
	var a Base
	if err := json.Unmarshal([]byte(t), &a); err != nil {
		log.Fatal("Unable to decode data.", err)
	}
	//for _, item := range a.Items {
	//		fmt.Println(item.Media.URL)
	//}

	const flickr = `
  <!doctype html>
  <body>
  {{range .Items}}
  <img src="{{.Media.URL}}">
  {{end}}
  </body>
  </html>
`
	templ := template.Must(template.New("xyz").Parse(flickr))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templ.Execute(w, a)
		if err != nil {
			log.Println("executing template:", err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
