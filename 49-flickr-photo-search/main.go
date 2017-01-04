/*https://api.flickr.com/services/feeds/photos_public.gne?format=json&tags=nature
  Create program with graphical interface that takes a string search and displays
  matching pictures.
Remove : jsonFlickrFeed(
*/

package main

import (
	"encoding/json"
	"fmt"
	"html"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
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

	//for _, item := range a.Items {
	//		fmt.Println(item.Media.URL)
	//}

	const flickr = `
  <!doctype html>
  <body>
  <form action="/search" method="post">
    <input type="text" name="tags" />
    <button type="submit"> submit </button>
  </form>
  {{range .Items}}
  <img src="{{.Media.URL}}">
  {{end}}
  </body>
  </html>
`
	templ := template.Must(template.New("xyz").Parse(flickr))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var t Base
		if r.Method == "POST" {
			tags := r.FormValue("tags")
			var err error
			t, err = getFlickrData(tags)
			if err != nil {
				log.Println("Error getting Flickr data", err)
				return
			}
		}

		err := templ.Execute(w, t)
		if err != nil {
			log.Println("executing template:", err)
		}
	})
	http.ListenAndServe(":8080", nil)
}
func getFlickrData(tags string) (Base, error) {
	v := url.Values{}
	v.Set("format", "json")
	v.Set("tags", tags)
	resp, err := http.Get("https://api.flickr.com/services/feeds/photos_public.gne?" + v.Encode())
	if err != nil {
		return Base{}, fmt.Errorf("Error getting URL: %s", err)
	}
	defer resp.Body.Close()

	b, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return Base{}, fmt.Errorf("Error reading response body: %s", err)
	}
	body := html.UnescapeString(string(b))
	t := strings.TrimPrefix(body, "jsonFlickrFeed(")
	t = strings.TrimSuffix(t, ")")
	var a Base
	if err := json.Unmarshal([]byte(t), &a); err != nil {
		return Base{}, fmt.Errorf("Error decoding json: %s", err)
	}
	return a, nil
}
