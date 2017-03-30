package controller

import (
	// "fmt"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"html/template"
	// "myapp/module"
	"net/http"
)

func ShowArticleList(w http.ResponseWriter, r *http.Request) {
	t := template.Must(template.ParseFiles("templates/header.html"))
	t.ExecuteTemplate(w, "header", nil)
}
