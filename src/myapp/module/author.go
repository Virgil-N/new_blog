package module

// mongodb字段必须全部小写，而且不能用下划线
type Author struct {
	Id        int64    `json: "id"`
	Name      string   `json: "name"`
	Avator    string   `json: "avator"`
	Age       int32    `json: "age"`
	Gender    string   `json: "gender"`
	Password  string   `json: "password"`
	Favourite []string `json: "favourite"`
	// 日期格式直接放在模板上显示2016-12-12 20:24:53 +0800 CST,还是改成string类型
	// RegisterDate time.Time `json: "registerdate"`
	RegisterDate string `json: "registerdate"`
	ArticleCount int32  `json: "articlecount"`
}
