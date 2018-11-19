package server

import (
	"common/config"
	"encoding/json"
	"github.com/gobestsdk/gobase/httpserver"
	"github.com/gobestsdk/gobase/log"
	"model"
	"net/http"
)

var (
	s                     = httpserver.New()
	methodnotfounthandler = &MethodNotFoundHandler{}
)

type MethodNotFoundHandler struct {
}

func (this *MethodNotFoundHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	log.Error(log.Fields{"err": "handlernotfound", "url": req.URL})
	r := model.CommonResp{
		Message: "路由不存在",
	}
	rb, err := json.Marshal(r)
	if err != nil {
		log.Error(log.Fields{"err": err, "json": "User_Upsertresp"})
	}

	resp.Write(rb)
}

func Run() {

	s.SetPort(config.APPConfig.HttpPort)
	s.Run()
}

func Stop() {
	s.Stop()
}
