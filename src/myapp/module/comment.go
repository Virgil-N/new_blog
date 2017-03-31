package module

type Comment struct {
	Id         int64  `json: "id"`
	Username   string `json: "username"`
	Text       string `json: "text"`
	CreateDate string `json: "createdate"`
	Title      string `json: "title"`
}
