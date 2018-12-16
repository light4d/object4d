package service

import (
	"github.com/gobestsdk/gobase/log"

	"github.com/light4d/object4d/common/server"
	"github.com/light4d/object4d/dao"
	"github.com/light4d/object4d/model"
	"github.com/minio/minio-go"
	"io"
)

func SearchObject4d(filter map[string]interface{}) (result []model.Object4d, err error) {
	log.Info(log.Fields{
		"func":   "SearchObject4ds",
		"filter": filter,
	})
	db := dao.DB(server.APPConfig.Mysql)
	err = db.Table("object4d").Find(&result, filter).Error
	if err != nil {
		log.Warn(log.Fields{
			"Find":   "object4ds",
			"Result": result,
			"Err":    err.Error(),
		})
	}
	log.Info(log.Fields{
		"func":   "GetObject4ds",
		"result": result,
	})
	return

}
func FcreateObject4d(recommendcon model.Miniocon, object model.Object4d, sourceobjectstream io.Reader) (n int64, err error) {
	log.Info(log.Fields{
		"func":     "CreateObject",
		"object4d": object,
	})

	db := dao.DB(server.APPConfig.Mysql)
	err = db.Table("object4d").Create(&object).Error
	if err != nil {
		log.Warn(log.Fields{
			"object":       object,
			"CreateObject": "DB",
			"Err":          err.Error(),
		})
		return
	}

	objects, err := SearchObject4d(map[string]interface{}{
		"t":   object.T,
		"lng": object.Lng,
		"lat": object.Lat,
	})
	if err != nil {
		return
	}
	if len(objects) == 0 {
		return 0, model.ErrLenNotEqual1
	}
	if len(objects) > 1 {
		return 0, model.ErrLenBigThan1
	}

	mc, err := dao.NewMinioclient(&recommendcon)
	if err != nil {
		log.Warn(log.Fields{
			"object4d":       object,
			"NewMinioclient": "NewMinioclient",
			"Err":            err.Error(),
		})
		return
	}
	err = mc.MakeBucket(object.Bucket(), "")
	if err != nil {
		log.Warn(log.Fields{
			"object4d":   object,
			"MakeBucket": "MakeBucket",
			"Err":        err.Error(),
		})
		return
	}
	n, err = mc.PutObject(object.Bucket(), object.Objectname(), sourceobjectstream, -1, minio.PutObjectOptions{})
	if err != nil {
		log.Warn(log.Fields{
			"object4d":     object,
			"CreateObject": "minio.PutObject",
			"Err":          err.Error(),
		})

	}
	return
}

func FgetObject(object model.Object4d) (stream io.Reader, err error) {
	objects, err := SearchObject4d(map[string]interface{}{
		"t":   object.T,
		"lng": object.Lng,
		"lat": object.Lat,
	})
	if err != nil {
		return
	}
	if len(objects) == 0 {
		return nil, model.ErrLenNotEqual1
	}
	if len(objects) > 1 {
		return nil, model.ErrLenBigThan1
	}

	mc, err := dao.NewMinioclientByid(object.M)
	if err != nil {
		return nil, err
	}
	e, err := mc.BucketExists(object.Bucket())
	if err != nil || !e {
		return nil, err
	}

	o, err := mc.GetObject(object.Bucket(), object.Objectname(), minio.GetObjectOptions{})
	if err != nil {
		return nil, err
	}

	return o, nil
}
