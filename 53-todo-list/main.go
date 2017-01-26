package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
)

type Todo struct {
	Name string
}

type Todos []Todo

func main() {

	t := Todos{}

	f, err := os.Open("data.json")
	if err != nil {
		log.Fatal("Unable to open file", err)
	}

	if err := json.NewDecoder(f).Decode(&t); err != nil {
		log.Fatal("Unable to decode.", err)
	}
	f.Close()

	templ := template.Must(template.New("xyz").Parse(form))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		err := templ.Execute(w, t)
		if err != nil {
			log.Println("executing template:", err)
		}
	})

	http.HandleFunc("/add", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			// Parse the form for Todo items Name
			name := r.FormValue("todo")

			if name != "" {
				// Create new instance of Todo struct set to Name
				todo := Todo{Name: name}
				// Append instance to t type Todos
				t = append(t, todo)
			}
		}
		// Redirect response to root URL "/"
		http.Redirect(w, r, "/", http.StatusFound)
	})

	//New handlefunc route = "/remove"
	http.HandleFunc("/remove", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			delete := r.FormValue("todo")
			for i, todo := range t {
				if todo.Name == delete {
					t = append(t[:i], t[i+1:]...)
				}
			}
		}
		http.Redirect(w, r, "/", http.StatusFound)

	})
	go http.ListenAndServe(":8080", nil)

	fmt.Println("Started Server")
	stop := make(chan os.Signal)
	signal.Notify(stop, os.Interrupt)
	<-stop
	fmt.Println("Shutting down.")

	f, err = os.Create("data.json")
	if err != nil {
		log.Fatal("Unable to open file", err)
	}
	defer f.Close()

	if err := json.NewEncoder(f).Encode(t); err != nil {
		log.Fatal("Unable to encode.", err)
	}
}

const form = `
<!doctype html>
<html>
<body>
<form action="/add" method="post">
  <input type="text" name="todo" />
  <button type="submit"> Add </button>
</form>
<form action="/remove" method="post">
  <input type="text" name="todo" />
  <button type="submit"> Remove </button>
</form>
<ul>
  {{range $}}
  <li>{{.Name}}</li>
  {{end}}
</ul>
</body>
</html>
`
