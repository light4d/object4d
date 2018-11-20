package service

import (
	"errors"
	"fmt"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/yourfs/dao"
	"github.com/light4d/yourfs/model"
)

func AddGroupusers(who, group string, us []string) (err error) {

	log.Info(log.Fields{
		"func":   "Addusers",
		"detail": who + " add users: " + fmt.Sprint(us) + " to " + group,
	})

	gs, err := SearchGroup(map[string]interface{}{
		"id": group,
	})
	if err != nil {
		return
	}
	if len(gs) == 0 {
		return model.ErrLenNotEqual1
	}
	if len(gs) > 1 {
		return model.ErrLenBigThan1
	}
	if gs[0].Parent != who {
		return errors.New("you no permission delete this group")
	}

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

	log.Info(log.Fields{
		"func":    "CreateGroup",
		"groupid": gs[0].ID,
	})
	return
}

func Deleteusers(who, group string, us []string) (err error) {
	log.Info(log.Fields{
		"func":   "Addusers",
		"detail": who + " delete users: " + fmt.Sprint(us) + " to " + group,
	})
	return
}
func Resetusers(who, group string, us []string) (err error) {
	log.Info(log.Fields{
		"func":   "Addusers",
		"detail": who + " reset users: " + fmt.Sprint(us) + " to " + group,
	})
	return
}
