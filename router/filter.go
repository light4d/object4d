package router

import (
	moehttp "github.com/light4d/yourfs/common/http"
	"github.com/light4d/yourfs/model"
	"github.com/light4d/yourfs/service"
	"github.com/qiniu/x/errors.v7"
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
		moehttp.Endresp(result, resp)
		return
	}
	if service.Checktoken(c.Value) == "" {
		result.Code = -2
		result.Error = errors.New("need token")
		moehttp.Endresp(result, resp)
		return
	} else {
		switch req.URL.Path {
		case "/user":
			user(resp, req)
		case "/group":
			group(resp, req)
		}
		return
	}
}
