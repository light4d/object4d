package router

import (
	"github.com/gobestsdk/gobase/httpserver"
	"github.com/light4d/lightlocation"
	"net/http"
)

func object4d(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		object4d_get(resp, req)
	case http.MethodPost:
		object4d_post(resp, req)
	default:
		httpserver.Options(req, resp)
	}
}

func object4d_get(resp http.ResponseWriter, req *http.Request) {
	//uid:=getuid(req)

	//service.GetObject(uid)
}
func object4d_post(resp http.ResponseWriter, req *http.Request) {
	lightlocation.GetLocation(req)
}
