package router

import (
	"github.com/light4d/yourfs/model"
	"github.com/light4d/yourfs/service"

	"errors"
	"github.com/gobestsdk/gobase/httpserver"
	"net/http"
)

func getuid(req *http.Request) string {
	c, err := req.Cookie("token")
	if err != nil {
		return ""
	}
	return service.Checktoken(c.Value)
}

func checktoken(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")

	result := model.CommonResp{}
	c, err := req.Cookie("token")
	if err != nil {
		result.Code = -2
		result.Error = "nedd token"
		httpserver.Endresp(result, resp)
		return
	}
	if service.Checktoken(c.Value) == "" {
		result.Code = -2
		result.Error = errors.New("need token")
		httpserver.Endresp(result, resp)
		return
	} else {
		switch req.URL.Path {
		case "/user":
			user(resp, req)
		case "/group":
			group(resp, req)
		case "/group/owner":
			group_setowner(resp, req)
		case "/group/user":
			group_user(resp, req)
		}
		return
	}
}
