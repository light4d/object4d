package service

import (
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/yourfs/dao"
	"github.com/light4d/yourfs/model"
	"io"
	"strconv"
	"time"
)

func Login(userid, password string) (string, error) {
	existuser, err := GetUser(userid)

	if err != nil {
		log.Info(log.Fields{
			"loginuser": existuser,
			"err":       err.Error(),
		})
		return "", err
	}
	if existuser == nil {
		return "", errors.New("用户未注册")

	}
	fmt.Println(existuser.Password, model.DBPassword(password))
	if (existuser.Password) != model.DBPassword(password) {
		return "", errors.New("用户名密码错误")
	}

	t := Token()

	dao.Redis().Set(t, userid, time.Hour*2)
	return t, nil
}

func Token() string {
	crutime := time.Now().Unix()
	h := md5.New()
	io.WriteString(h, strconv.FormatInt(crutime, 10))
	token := fmt.Sprintf("%x", h.Sum(nil))
	return token
}
