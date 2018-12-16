package dao

import (
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/common/server"
	"github.com/light4d/object4d/model"
	"github.com/minio/minio-go"
)

func GetMinioconfig(minioid int) (m *model.Miniocon, err error) {
	m = new(model.Miniocon)
	err = DB(server.APPConfig.Mysql).Table("miniocon").Where("id = ?", minioid).Find(m).Error
	return
}

func NewMinioclient(cfg *model.Miniocon) (c *minio.Client, err error) {
	c, err = minio.New(cfg.Endpoint, cfg.Ak, cfg.Sk, cfg.Secure)
	if err != nil {
		log.Warn(log.Fields{
			"Func": "NewMinioclient",
			"Err":  err.Error(),
		})
	}
	log.Info(log.Fields{
		"minio": cfg,
	})
	return
}
func NewMinioclientByid(minioid int) (c *minio.Client, err error) {
	// 初使化minio client对象。
	cfg, err := GetMinioconfig(minioid)
	if err != nil {
		log.Warn(log.Fields{
			"Func":   "GetMinioconfig",
			"Err":    err.Error(),
			"Detail": cfg,
		})
		return nil, err
	}
	return NewMinioclient(cfg)
}
