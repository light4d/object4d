package router

import (
	moehttp "github.com/light4d/yourfs/common/http"
	"net/http"
)

func object(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		object_get(resp, req)
	case http.MethodPost:
		object_post(resp, req)
	case http.MethodDelete:
		object_delete(resp, req)
	default:
		moehttp.Options(req, resp)
	}
}

func object_get(resp http.ResponseWriter, req *http.Request) {

}
func object_post(resp http.ResponseWriter, req *http.Request) {

}
func object_delete(resp http.ResponseWriter, req *http.Request) {

}
