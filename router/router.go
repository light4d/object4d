package router

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/user", checktoken)
	http.HandleFunc("/group", checktoken)
	http.HandleFunc("/group/owner", checktoken)
	http.HandleFunc("/group/user", checktoken)
	http.HandleFunc("/login", login)
	http.HandleFunc("/", object)

}
