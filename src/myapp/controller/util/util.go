package util

import (
	"fmt"
	"net/http"
)

func IsLogin(w http.ResponseWriter, r *http.Request) bool {
	_, err := r.Cookie("authorId")
	if err != nil {
		fmt.Println("请登录!")
		// http.Redirect(w, r, "/showLogin", 302)
		return false
	}
	return true
}
