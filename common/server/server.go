package server

import (
	"github.com/gobestsdk/gobase/httpserver"
)

const Appname = "ojbect4d"

var (
	O = httpserver.New(Appname)
)

func Run() {

	O.SetPort(APPConfig.Object4dPort)
	O.Run()

}

func Stop() {
	O.Stop()
}
