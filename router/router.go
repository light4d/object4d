package router

import (
	"net/http"
)

func Init() {

	http.HandleFunc("/user", user)
}
