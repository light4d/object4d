package service

import (
	"errors"
	"fmt"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/yourfs/dao"
	"github.com/light4d/yourfs/model"
	"time"
)

func SearchGroupuser(filter map[string]interface{}) (result []model.Groupuser, err error) {
	log.Info(log.Fields{
		"func":   "SearchGroupuser",
		"filter": filter,
	})

	db := dao.DB()
	err = db.Table("groupuser").Find(&result, filter).Error
	if err != nil {
		log.Warn(log.Fields{
			"Find":   "groupuser",
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
func AddGroupusers(who, group string, us []string) (err error) {

	log.Info(log.Fields{
		"func":   "Addusers",
		"detail": who + " add users: " + fmt.Sprint(us) + " to " + group,
	})

	g, ex := CheckGroupExist(group)
	if !ex {
		return errors.New("group not found")
	}
	if g.Parent != who {
		return errors.New("you no permission add user to this group")
	}

	for _, u := range us {
		_, ex := CheckUserExist(u)
		if !ex {
			return model.NewErrData("user not exist", u)
		}
	}
	db := dao.DB()
	for _, u := range us {
		err = db.Table("groupuser").Create(&model.Groupuser{
			ID:       group,
			User:     u,
			Jointime: time.Now(),
		}).Error
		if err != nil {
			log.Warn(log.Fields{
				"groupuser":       group,
				"CreateGroupuser": "DB",
				"Err":             err.Error(),
			})
			db.Rollback()
			return
		}
	}
	db.Commit()

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
