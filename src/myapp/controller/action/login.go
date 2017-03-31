package action

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"myapp/module"
	"net/http"
	"strconv"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {

	username := r.FormValue("username")
	password := r.FormValue("password")

	if username != "" && password != "" {
		session, err := mgo.Dial("localhost:27017")
		if err != nil {
			w.Write([]byte("连接数据库失败!"))
			return
		}
		defer session.Close()
		c := session.DB("new_blog").C("author")
		var result module.Author
		err = c.Find(bson.M{"name": username}).One(&result)
		if err != nil {
			w.Write([]byte("查找用户失败!"))
			// panic(err)
			return
		}
		if password != result.Password {
			w.Write([]byte("用户名或密码错误!"))
			return
		}
		if password == result.Password {
			//设置cookie，未对中文进行处理
			expires := time.Now().AddDate(0, 0, 1)
			userCookie := http.Cookie{
				Name:    "isAuthor",
				Value:   "yes",
				Path:    "/",
				Expires: expires,
				MaxAge:  86400,
			}
			http.SetCookie(w, &userCookie)

			authorCookie := http.Cookie{
				Name:    "authorId",
				Value:   strconv.Itoa(int(result.Id)),
				Path:    "/",
				Expires: expires,
				MaxAge:  86400,
			}
			http.SetCookie(w, &authorCookie)

			w.Write([]byte("登录成功!"))
			return
		}
	} else {
		w.Write([]byte("用户名及密码是必填项!"))
		return
	}

}
