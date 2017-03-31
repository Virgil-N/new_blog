package module

type Article struct {
	Id         int64    `json: "id"`
	Author     string   `json: "author"`
	Title      string   `json: "title"`
	Content    string   `json: "content"`
	TotalWords int64    `json: "totalwords"`
	CreateDate string   `json: "createdate"`
	ModifyDate string   `json: "modifydate"`
	Category   []string `json: "category"`
	CommentIds []int64  `json: "commentids"`
	CanComment bool     `json: "cancomment"`
	Read       int32    `json: "hasread"`
	Commented  int32    `json: "commented"`
}
