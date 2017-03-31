package page

import (
	"fmt"
	"html/template"
	"net/http"
)

func ShowLogin(w http.ResponseWriter, r *http.Request) {
	fmt.Println("...")
	t := template.Must(template.ParseFiles("templates/login.html", "templates/footer.html"))
	t.ExecuteTemplate(w, "login", nil)
}
