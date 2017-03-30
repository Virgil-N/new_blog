package main

import (
	"github.com/gorilla/mux"
	"log"
	"myapp/controller"
	"net/http"
)

func main() {
	r := mux.NewRouter()
	http.Handle("/css/", http.FileServer(http.Dir("public")))
	http.Handle("/js/", http.FileServer(http.Dir("public")))
	http.Handle("/vendor/", http.FileServer(http.Dir("public")))
	http.Handle("/img/", http.FileServer(http.Dir("public")))

	r.HandleFunc("/", controller.ShowArticleList)
	r.HandleFunc("/index", controller.ShowArticleList)

	http.Handle("/", r)

	if err := http.ListenAndServe("localhost:9090", nil); err != nil {
		log.Fatalln("ListenAndServe: ", err)
	}
}
