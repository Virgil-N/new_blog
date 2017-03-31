package page

import (
	// "fmt"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"html/template"
	"myapp/module"
	"net/http"
)

type ArticleAllInfo struct {
	Author   module.Author
	Articles []module.Article
	Comments []module.Comment
	IsAuthor bool
	Paging   Pagination
}

type Pagination struct {
	Total    int
	Page     int
	PageSize int
}

const PAGE = 1
const PAGE_SIZE = 10

func ShowArticleList(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/nav.html"))
	t.ExecuteTemplate(w, "nav", nil)
}
