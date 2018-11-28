package server

import (
	"github.com/gobestsdk/gobase/httpserver"

	"github.com/light4d/yourfs/common/config"
)

var (
	M = httpserver.New("yourfs")
	F = httpserver.New("fs")
)

func Run() {

	M.SetPort(config.APPConfig.HttpPort)
	F.SetPort(config.APPConfig.FsPort)
	go M.Run()
	F.Run()
}

func Stop() {
	M.Stop()
	F.Stop()
}
