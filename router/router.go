package router

import (
	"net/http"
)

func Init() {
	http.HandleFunc("/", checktoken)
	http.HandleFunc("/login", login)

}
