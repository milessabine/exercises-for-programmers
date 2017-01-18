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

	http.ListenAndServe(":8080", nil)
}

//New handlerfunc route = "/remove"{
//if Method == "DELETE"{
//delete:= r.FormValue(name)
//for t []Todos
//if t.Name == delete {
//	t = delete(t,delete)
//}
//Redirect response to root URL "/"
//create new form. set method to delete and action to "/remove"
const form = `
<!doctype html>
<html>
<body>
<form action="/add" method="post">
  <input type="text" name="todo" />
  <button type="submit"> Add </button>
</form>
<ul>
  {{range $}}
  <li>{{.Name}}</li>
  {{end}}
</ul>
</body>
</html>
`
