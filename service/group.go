package service

import (
	"fmt"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/yourfs/dao"
	"github.com/light4d/yourfs/model"
	"github.com/qiniu/x/errors.v7"
	"regexp"
	"time"
)

func checkandfixCreateGroup(user *model.Group) (err error) {
	//TODO your code
	m, err := regexp.Match("^[a-zA-Z0-9_-]{4,16}$", []byte(user.ID))
	if user.ID == "" || err != nil || !m {
		return model.NewErrData("group id error format", user.ID)
	}
	user.Registetime = time.Now()
	return
}
func SetOwner(who, newowner, groupid string) (err error) {
	log.Info(log.Fields{
		"func":   "SetOwner",
		"detail": who + " set group " + groupid + " to " + newowner,
	})

	filter := map[string]interface{}{
		"id":   groupid,
		"type": "group",
	}
	groups, err := SearchGroup(filter)
	if err != nil {
		log.Warn(log.Fields{
			"GetGroups": groups,
			"Err":       err.Error(),
		})
		return err
	}

	if len(groups) == 0 {
		return model.ErrLenEqual0
	}
	if len(groups) > 1 {
		return model.ErrLenBigThan1
	}
	if groups[0].Parent != who {
		return errors.New("you no permission set this group")
	}

	filter_user := map[string]interface{}{
		"id":   newowner,
		"type": "",
	}
	us, err := SearchUser(filter_user)
	if err != nil {
		log.Warn(log.Fields{
			"GetUsers": us,
			"Err":      err.Error(),
		})
		return err
	}

	if len(us) == 0 {
		return model.NewErr("newowner not found")
	}
	if len(us) > 1 {
		return model.ErrLenBigThan1
	}

	err = dao.DB().Where(filter).Table("user").Model(new(model.Group)).Update(
		map[string]interface{}{
			"parent": newowner,
		}).Error
	if err != nil {
		log.Warn(log.Fields{
			"Group.Update": groups,
			"Where":        filter,
			"Err":          err.Error(),
		})
	}
	return
}

func CreateGroup(parend string, group model.Group) (groupid string, err error) {
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
	err = db.Model(new(model.Group)).Create(&group).Error
	if err != nil {
		log.Warn(log.Fields{
			"user":       group,
			"CreateUser": "DB",
			"Err":        err.Error(),
		})
		return
	}

	gs, err := SearchGroup(map[string]interface{}{
		"name": group.Name,
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

func SearchGroup(filter map[string]interface{}) (result []model.Group, err error) {
	log.Info(log.Fields{
		"func":   "SearchGroup",
		"filter": filter,
	})

	filter["type"] = "group"

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
		"func":   "SearchGroup",
		"result": result,
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
	groups, err := SearchGroup(filter)
	if err != nil {
		log.Warn(log.Fields{
			"GetUsers": groups,
			"Err":      err.Error(),
		})
		return err
	}

	if len(groups) == 0 {
		return model.ErrLenEqual0
	}
	if len(groups) > 1 {
		return model.ErrLenBigThan1
	}
	if groups[0].Parent != who {
		return errors.New("you no permission delete this group")
	}

	err = dao.DB().Where(filter).Table("user").Model(new(model.Group)).Delete(&model.User{}).Error
	if err != nil {
		log.Warn(log.Fields{
			"Model.Delete": groups,
			"Where":        filter,
			"Err":          err.Error(),
		})
	}
	return err
}

var allowupdateGroup = map[string]interface{}{
	"name": "",
	"face": "",
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
	groups, err := SearchGroup(filter)
	if err != nil {
		log.Warn(log.Fields{
			"GetUsers": groups,
			"Err":      err.Error(),
		})
		return err
	}

	if len(groups) == 0 {
		return model.ErrLenEqual0
	}
	if len(groups) > 1 {
		return model.ErrLenBigThan1
	}
	if groups[0].Parent != who {
		return errors.New("you no permission delete this group")
	}

	db := dao.DB()
	err = db.Model(new(model.Group)).Table("user").Where("id = ?", groupid).Where("type = 'group'").Updates(updater).Error
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
