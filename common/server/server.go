package server

import (
	"github.com/gobestsdk/gobase/httpserver"

	"github.com/light4d/object4d/common/config"
)

var (
	O = httpserver.New("ojbect4d")
)

func Run() {

	O.SetPort(config.APPConfig.Object4dPort)
	O.Run()

}

func Stop() {
	O.Stop()
}
