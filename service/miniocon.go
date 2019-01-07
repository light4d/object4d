package service

import (
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/common/server"
	"github.com/light4d/object4d/dao"
	"github.com/light4d/object4d/model"
	"time"
)

func Minioconcount() (c []model.Miniocon) {
	db := dao.DB(server.APPConfig.Mysql)
	defer db.Close()
	err := db.Table("miniocon").Find(&c).Error
	if err != nil {
		log.Warn(log.Fields{
			"object":       "miniocon",
			"CreateObject": "DB",
			"Err":          err.Error(),
		})
	}

	return
}
func RendMinioconid() model.Miniocon {
	ms := Minioconcount()
	log.Info(log.Fields{
		"minio con": ms,
	})
	idx := time.Now().Second() % len(Minioconcount())
	log.Info(log.Fields{
		"index": idx,
	})
	return ms[idx]
}
