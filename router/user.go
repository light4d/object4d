package router

import (
	"encoding/json"
	"errors"
	"io/ioutil"

	moehttp "github.com/light4d/yourfs/common/http"
	"github.com/light4d/yourfs/model"
	"net/http"
	"strconv"

	"github.com/light4d/yourfs/service"
)

func user(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		user_get(resp, req)
	case http.MethodPost:
		user_post(resp, req)
	case http.MethodPut:
		user_put(resp, req)
	case http.MethodDelete:
		user_delete(resp, req)
	default:
		moehttp.Options(req, resp)
	}
}
func user_get(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}
	filter := moehttp.Getfilter(req)
	result.Result, result.Error = service.SearchUsers(filter)
	moehttp.Endresp(result, resp)
}
func user_post(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	user := model.User{}
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		result.Error = err
		moehttp.Endresp(result, resp)
		return
	}
	err = json.Unmarshal(body, &user)
	if err != nil {
		result.Error = err
		moehttp.Endresp(result, resp)
		return
	}

	userid, err := service.CreateUser(user)
	if err != nil {
		result.Error = err
		moehttp.Endresp(result, resp)
		return
	}
	result.Result = struct {
		UserID int
	}{UserID: userid}
	moehttp.Endresp(result, resp)
}

func user_put(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	id, err := strconv.Atoi(req.URL.Query().Get("id"))
	if err != nil {
		result.Error = err
		moehttp.Endresp(result, resp)
		return
	}
	updater := make(map[string]interface{})
	err = moehttp.Unmarshalreqbody(req, &updater)
	if err != nil {
		result.Error = err
		moehttp.Endresp(result, resp)
		return
	}
	result.Error = service.UpdateUser(id, updater)
	moehttp.Endresp(result, resp)
}
func user_delete(resp http.ResponseWriter, req *http.Request) {
	result := model.CommonResp{}

	useridstr := req.URL.Query().Get("id")

	if useridstr != "" {
		userid, err := strconv.Atoi(useridstr)
		if err != nil {
			result.Error = model.NewErr("err = strconv.Atoi(id)")
			moehttp.Endresp(result, resp)
			return
		}

		result.Error = service.DeleteUser(userid)

		moehttp.Endresp(result, resp)
		return
	} else {
		result.Error = errors.New("whick one do you want to delete?")
		moehttp.Endresp(result, resp)
		return
	}

}
