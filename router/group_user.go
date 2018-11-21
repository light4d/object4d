package router

import (
	"encoding/json"
	"errors"
	moehttp "github.com/light4d/yourfs/common/http"
	"github.com/light4d/yourfs/model"
	"github.com/light4d/yourfs/service"
	"io/ioutil"
	"net/http"
)

func group_user(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		group_user_get(resp, req)
	case http.MethodPost:
		group_user_post(resp, req)
	//case http.MethodPut:
	//	group_user_put(resp, req)
	//case http.MethodDelete:
	//	group_user_delete(resp, req)
	default:
		moehttp.Options(req, resp)
	}
}

func group_user_get(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	filter := moehttp.Getfilter(req)
	gs, err := service.SearchGroupuser(filter)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
	} else {
		result.Result = gs
	}
	moehttp.Endresp(result, resp)
}
func group_user_post(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	us := make([]string, 0)
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		moehttp.Endresp(result, resp)
		return
	}
	err = json.Unmarshal(body, &us)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		moehttp.Endresp(result, resp)
		return
	}
	id := req.URL.Query().Get("id")
	if id == "" {
		result.Code = -1
		result.Error = errors.New("id不能为空")
		moehttp.Endresp(result, resp)
		return
	}
	uid := getuid(req)
	err = service.AddGroupusers(uid, id, us)
	if err != nil {
		result.Code = -1
		result.Error = err.Error()
		moehttp.Endresp(result, resp)
		return
	}

	moehttp.Endresp(result, resp)
}
