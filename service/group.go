package service

import (
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/yourfs/dao"
	"github.com/light4d/yourfs/model"
	"github.com/qiniu/x/errors.v7"
	"regexp"
	"time"
)

func checkandfixCreateGroup(user *model.User) (err error) {
	//TODO your code

	m, err := regexp.Match("^[a-zA-Z0-9_-]{4,16}$", []byte(user.ID))
	if user.ID == "" || err != nil || !m {
		return model.NewErrData("group id error format", user.ID)
	}

	user.Registetime = time.Now()
	return
}
func CreateGroup(parend string, group model.User) (groupid string, err error) {
	log.Info(log.Fields{
		"func": "CreateUser",
		"user": group,
		"who":  parend,
	})

	err = checkandfixCreateGroup(&group)
	if err != nil {
		return
	}
	group.Type = "group"
	group.Parent = parend
	db := dao.DB()
	err = db.Table("user").Create(&group).Error
	if err != nil {
		log.Warn(log.Fields{
			"user":       group,
			"CreateUser": "DB",
			"Err":        err.Error(),
		})
		return
	}

	gs, err := SearchUserorGroup(map[string]interface{}{
		"name": group.Name,
		"type": "group",
	})
	if err != nil {
		return
	}
	if len(gs) == 0 {
		return "", model.ErrLenNotEqual1
	}
	if len(gs) > 1 {
		return "", model.ErrLenBigThan1
	}
	groupid = gs[0].ID
	log.Info(log.Fields{
		"func":    "CreateGroup",
		"groupid": gs[0].ID,
	})
	return
}

func DeleteGroup(who string, groupid string) (err error) {
	log.Info(log.Fields{
		"func":    "DeleteGroup",
		"groupid": groupid,
		"who":     who,
	})
	filter := map[string]interface{}{
		"id":   groupid,
		"type": "group",
	}
	users, err := SearchUserorGroup(filter)
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
	if users[0].Parent != who {
		return errors.New("you no permission delete this group")
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

var allowupdateGroup = map[string]interface{}{
	"name":   "",
	"parend": "",
	"face":   "",
}

func UpdateGroup(who, groupid string, updater map[string]interface{}) (err error) {
	log.Info(log.Fields{
		"func":    "UpdateUser",
		"id":      groupid,
		"who":     who,
		"updater": updater,
	})

	errfields := checkupdate(updater, allowupdateGroup)
	if len(errfields) > 0 {
		return model.NewErrData(model.FieldCannotupdate, errfields)
	}

	filter := map[string]interface{}{
		"id":   groupid,
		"type": "group",
	}
	users, err := SearchUserorGroup(filter)
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
	if users[0].Parent != who {
		return errors.New("you no permission delete this group")
	}

	db := dao.DB()
	err = db.Table("user").Where("id = ?", groupid).Where("type = 'group'").Updates(updater).Error
	if err != nil {
		log.Warn(log.Fields{
			"func":    "UpdateUser Updates",
			"id":      groupid,
			"updater": updater,
			"Err":     err.Error(),
		})
	}
	return
}
