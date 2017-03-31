package page

import (
	"html/template"
	"net/http"
)

func ShowLogin(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/login.html"))
	t.ExecuteTemplate(w, "login", nil)
}
