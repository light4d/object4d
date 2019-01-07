package router

import (
	"github.com/gobestsdk/gobase/httpserver"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/common"
	"github.com/light4d/object4d/common/server"
	"github.com/light4d/object4d/model"
	"github.com/light4d/object4d/service"
	"io/ioutil"
	"net/http"
)

func Object4d(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-Type", "application/json")
	switch req.Method {
	case http.MethodGet:
		object4d_get(resp, req)
	case http.MethodPost:
		object4d_post(resp, req)
	default:
		httpserver.Options(req, resp, "application/octet-stream", server.Appname, AccessControlAllowMethods())
	}
}

func object4d_get(resp http.ResponseWriter, req *http.Request) {
	log.Info(log.Fields{
		"uri":    req.RequestURI,
		"method": req.Method,
	})
	object4d := model.ParseObject4d(req.RequestURI)
	log.Info(log.Fields{
		"object4d": object4d,
	})
	o, err := service.FgetObject(*object4d)
	if err != nil {
		result := model.CommonResp{
			Error: err.Error(),
			Code:  -1,
		}
		log.Warn(log.Fields{
			"func": "FgetObject",
			"Err":  err.Error(),
		})
		Endresp(result, resp)
		return
	}
	ost, err := o.Stat()
	if err != nil {
		result := model.CommonResp{
			Error: err.Error(),
			Code:  -1,
		}
		log.Warn(log.Fields{
			"func": "Stat",
			"Err":  err.Error(),
		})
		Endresp(result, resp)
		return
	}
	bs, err := ioutil.ReadAll(o)
	if err != nil {
		result := model.CommonResp{
			Error: err.Error(),
			Code:  -1,
		}
		log.Warn(log.Fields{
			"func": "ReadAll",
			"Err":  err.Error(),
		})
		Endresp(result, resp)
		return
	}
	resp.Header().Set("Content-type", ost.ContentType)
	resp.Header().Set("Server", server.Appname)
	resp.Write(bs)
}
func object4d_post(resp http.ResponseWriter, req *http.Request) {
	obj4d, err := service.GetLocation(req)
	if err != nil {
		result := model.CommonResp{
			Error: err.Error(),
			Code:  -1,
		}
		log.Warn(log.Fields{
			"error": err.Error(),
		})
		Endresp(result, resp)
		return
	}

	recommendcon := service.RendMinioconid()
	object4d := &model.Object4d{
		Lng: obj4d.Lng,
		Lat: obj4d.Lat,
		T:   obj4d.T,
		M:   recommendcon.ID,
	}
	contentType := req.Header.Get(common.Ctype)
	n, err := service.FcreateObject4d(recommendcon, *object4d, req.Body, contentType)
	if err != nil {
		result := model.CommonResp{
			Error: err.Error(),
			Code:  -1,
		}
		log.Warn(log.Fields{
			"func": "FcreateObject4d",
			"Err":  err.Error(),
		})
		Endresp(result, resp)
		return
	} else {
		result := model.CommonResp{
			Result: map[string]interface{}{
				"object4d": object4d,
				"url":      object4d.Url(),
				"size":     n,
			},
			Code: 0,
		}

		Endresp(result, resp)
		return
	}

}
