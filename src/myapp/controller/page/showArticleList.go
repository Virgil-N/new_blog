package page

import (
	// "fmt"
	// "gopkg.in/mgo.v2"
	// "gopkg.in/mgo.v2/bson"
	"html/template"
	// "math"
	// "myapp/module"
	"net/http"
	// "strconv"
	// "time"
)

// type ArticleAllInfo struct {
// 	Author   module.Author
// 	Articles []module.Article
// 	Comments []module.Comment
// 	IsAuthor string
// 	Paging   Pagination
// }

// type Pagination struct {
// 	Total    int
// 	Page     int
// 	PageSize int
// }

// const PAGE = 1
// const PAGE_SIZE = 10

func ShowArticleList(w http.ResponseWriter, r *http.Request) {

	t := template.Must(template.ParseFiles("templates/articleList.html", "templates/components/header.html", "templates/components/footer.html", "templates/components/info.html"))
	t.ExecuteTemplate(w, "articleList", nil)
}
