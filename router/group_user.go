package router

import (
	moehttp "github.com/light4d/yourfs/common/http"
	"github.com/light4d/yourfs/model"
	"github.com/light4d/yourfs/service"
	"net/http"
)

func group_user(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		group_user_get(resp, req)
	case http.MethodPost:
		group_user_post(resp, req)
	case http.MethodPut:
		group_user_put(resp, req)
	case http.MethodDelete:
		group_user_delete(resp, req)
	default:
		moehttp.Options(req, resp)
	}
}

func group_user_get(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	filter := moehttp.Getfilter(req)
	gs, err := service.SearchUser(filter)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
	} else {
		result.Result = gs
	}
	moehttp.Endresp(result, resp)
}
