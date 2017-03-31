package main

import (
	"github.com/gorilla/mux"
	"log"
	"myapp/controller/page"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/vendor/", http.FileServer(http.Dir("public")))
	http.Handle("/img/", http.FileServer(http.Dir("public")))

	r.HandleFunc("/", page.ShowArticleList)
	r.HandleFunc("/index", page.ShowArticleList)
	r.HandleFunc("/login", page.ShowLogin)

	http.Handle("/", r)

	if err := http.ListenAndServe("localhost:9090", nil); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
