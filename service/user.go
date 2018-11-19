package service

import (
	"errors"
	"github.com/gobestsdk/gobase/log"

	"github.com/light4d/yourfs/dao"
	"github.com/light4d/yourfs/model"
	"time"
)

var allowupdateUser = map[string]interface{}{
	"birth":        "",
	"name":         "",
	"password":     "",
	"gender":       "",
	"face":         "",
	"introduction": "",
}

func checkandfixCreateUser(user *model.User) (err error) {
	//TODO your code
	user.RegisteTime = time.Now()
	return
}
func SearchUsers(filter map[string]interface{}) (result []model.User, err error) {
	log.Info(log.Fields{
		"func":   "SearchUsers",
		"filter": filter,
	})
	db := dao.DB()

	err = db.Model(&model.User{}).Find(&result, filter).Error
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
	err = dao.DB().Where(filter).Delete(&model.User{}).Error
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
	err = db.Model(&model.User{}).Create(&user).Error
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
	err = db.Model(&model.User{}).Where("id = ?", id).Updates(updater).Error
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
	err := db.Model(new(model.User)).Where("id = ?", userid).Find(&users).Error
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
func Login(userid, password string) error {
	existuser, err := GetUser(userid)
	log.Info(log.Fields{
		"loginuser": existuser,
		"err":       err.Error(),
	})
	if err != nil {

		return err
	}
	if existuser == nil {
		return errors.New("用户未注册")

	}
	if (existuser.Password) != model.DBPassword(password) {
		return errors.New("用户名密码错误")

	}
	return nil
}
