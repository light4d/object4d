package router

import (
	"github.com/gobestsdk/gobase/httpserver"
	"github.com/light4d/lightlocation"
	"github.com/light4d/object4d/model"
	"github.com/light4d/object4d/service"
	"io/ioutil"
	"net/http"
	"time"
)

func object4d(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		object4d_get(resp, req)
	case http.MethodPost:
		object4d_post(resp, req)
	default:
		httpserver.Options(req, resp, "application/octet-stream", AccessControlAllowMethods())
	}
}

func object4d_get(resp http.ResponseWriter, req *http.Request) {
	object4d := model.ParseObject4d(req.URL.RawQuery)
	r, err := service.FgetObject(object4d)

	if err != nil {
		result := model.CommonResp{
			Error: err.Error(),
			Code:  -1,
		}
		Endresp(result, resp)
	}
	bs, err := ioutil.ReadAll(r)
	if err != nil {
		result := model.CommonResp{
			Error: err.Error(),
			Code:  -1,
		}
		Endresp(result, resp)
	}
	resp.Write(bs)
}
func object4d_post(resp http.ResponseWriter, req *http.Request) {
	lng, lat, err := lightlocation.GetLocation(req)

	if err != nil {
		result := model.CommonResp{
			Error: err.Error(),
			Code:  -1,
		}
		Endresp(result, resp)
	}

	object4d := model.Object4d{
		Lng: lng,
		Lat: lat,
		T:   time.Now().Format("2006-01-02 15:04:05"),
		M:   service.RendMinioconid(),
	}
	n, err := service.FcreateObject4d(object4d, req.Body)
	if err != nil {
		result := model.CommonResp{
			Error: err.Error(),
			Code:  -1,
		}
		Endresp(result, resp)
	} else {
		result := model.CommonResp{
			Result: n,
			Code:   0,
		}
		Endresp(result, resp)
	}

}
