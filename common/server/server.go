package server

import (
	"github.com/gobestsdk/gobase/httpserver"
)

var (
	O = httpserver.New("ojbect4d")
)

func Run() {

	O.SetPort(APPConfig.Object4dPort)
	O.Run()

}

func Stop() {
	O.Stop()
}
