package service

import (
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/dao"
	"time"
)

func Minioconcount() (c []int) {
	db := dao.DB()
	err := db.Table("miniocon").Find(&c).Select("id").Error
	if err != nil {
		log.Warn(log.Fields{
			"object":       "miniocon",
			"CreateObject": "DB",
			"Err":          err.Error(),
		})
	}
	return
}
func RendMinioconid() int {
	return Minioconcount()[time.Now().Second()%len(Minioconcount())]
}
