package service

import (
	"github.com/gobestsdk/gobase/log"

	"github.com/light4d/yourfs/dao"
	"github.com/light4d/yourfs/model"
	"regexp"
	"time"
)

var allowupdateUser = map[string]interface{}{
	"name":     "",
	"password": "",
	"face":     "",
}

func checkandfixCreateUser(user *model.User) (err error) {
	//TODO your code

	m, err := regexp.Match("^[a-zA-Z0-9_-]{4,16}$", []byte(user.ID))
	if user.ID == "" || err != nil || !m {
		return model.NewErrData("用户名格式不对", user.ID)
	}
	user.Password = model.DBPassword(user.Password)
	user.Registetime = time.Now()
	return
}
func SearchUsers(filter map[string]interface{}) (result []model.User, err error) {
	log.Info(log.Fields{
		"func":   "SearchUsers",
		"filter": filter,
	})
	db := dao.DB()

	err = db.Table("user").Find(&result, filter).Error
	if err != nil {
		log.Warn(log.Fields{
			"Find":   "users",
			"Result": result,
			"Err":    err.Error(),
		})
	}
	for i, _ := range result {
		result[i] = *(result[i].FixShow())
	}
	log.Info(log.Fields{
		"func":   "GetUsers",
		"result": result,
	})
	return

}

func DeleteUser(userid string) (err error) {
	log.Info(log.Fields{
		"func":   "DeleteUser",
		"userid": userid,
	})
	filter := map[string]interface{}{
		"id": userid,
	}
	users, err := SearchUsers(filter)

	if err != nil {
		log.Warn(log.Fields{
			"GetUsers": users,
			"Err":      err.Error(),
		})
		return err
	}

	if len(users) == 0 {
		return model.ErrLenEqual0
	}
	if len(users) > 1 {
		return model.ErrLenBigThan1
	}
	err = dao.DB().Where(filter).Table("user").Delete(&model.User{}).Error
	if err != nil {
		log.Warn(log.Fields{
			"Model.Delete": users,
			"Where":        filter,
			"Err":          err.Error(),
		})
	}
	return err
}
func CreateUser(user model.User) (userid string, err error) {
	log.Info(log.Fields{
		"func": "CreateUser",
		"user": user,
	})

	err = checkandfixCreateUser(&user)
	if err != nil {
		return
	}

	db := dao.DB()
	err = db.Table("user").Create(&user).Error
	if err != nil {
		log.Warn(log.Fields{
			"user":       user,
			"CreateUser": "DB",
			"Err":        err.Error(),
		})
		return
	}

	users, err := SearchUsers(map[string]interface{}{
		"name": user.Name,
	})
	if err != nil {
		return
	}
	if len(users) == 0 {
		return "", model.ErrLenNotEqual1
	}
	if len(users) > 1 {
		return "", model.ErrLenBigThan1
	}
	userid = (users[0].ID)
	log.Info(log.Fields{
		"func":   "CreateUser",
		"userid": userid,
	})
	return
}
func checkupdateUser(updater map[string]interface{}) (disableupdatefields []string) {
	disableupdatefields = make([]string, 0)
	for k, _ := range updater {
		if _, have := allowupdateUser[k]; !have {
			delete(updater, k)
			disableupdatefields = append(disableupdatefields, k)
		}
	}
	return
}
func UpdateUser(id string, updater map[string]interface{}) (err error) {
	log.Info(log.Fields{
		"func":    "UpdateUser",
		"id":      id,
		"updater": updater,
	})

	errfields := checkupdateUser(updater)
	if len(errfields) > 0 {
		return model.NewErrData(model.FieldCannotupdate, errfields)
	}
	db := dao.DB()
	err = db.Table("user").Where("id = ?", id).Updates(updater).Error
	if err != nil {
		log.Warn(log.Fields{
			"func":    "UpdateUser Updates",
			"id":      id,
			"updater": updater,
			"Err":     err.Error(),
		})
	}
	return
}

func GetUser(userid string) (*model.User, error) {
	db := dao.DB()
	users := make([]model.User, 0)
	err := db.Table("user").Where("id = ?", userid).Find(&users).Error
	if err != nil {
		return nil, err
	}
	if len(users) == 0 {
		return nil, nil
	}
	if len(users) == 1 {
		return &(users[0]), nil
	} else {
		return nil, model.ErrLenBigThan1
	}

}
