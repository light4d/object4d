package router

import (
	"github.com/light4d/object4d/model"

	"encoding/json"

	"github.com/gobestsdk/gobase/httpserver"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/common/server"
	"net/http"
	"strings"
)

func Endresp(result model.CommonResp, resp http.ResponseWriter) {
	log.Info(log.Fields{
		"resp": result,
	})
	httpserver.Header(resp, "application/json", server.Appname, AccessControlAllowMethods())

	r, _ := json.Marshal(result)
	resp.Write(r)
}
func AccessControlAllowMethods() string {
	var method = []string{
		http.MethodGet,
		http.MethodPost,
	}
	return strings.Join(method, ",")
}
