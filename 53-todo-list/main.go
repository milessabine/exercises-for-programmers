package main

import (
	"html/template"
	"log"
	"net/http"
)

type Todo struct {
	Name string
}

type Todos []Todo

func main() {

	var t Todos

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
	http.ListenAndServe(":8080", nil)

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
