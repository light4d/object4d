package router

import (
	"github.com/light4d/object4d/common/server"
)

func Init() {
	server.O.ServerMux.HandleFunc("/", Object4d)
}
