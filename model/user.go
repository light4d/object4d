package model

import (
	"crypto/md5"
	"encoding/hex"
	"github.com/gobestsdk/gobase/utils"
	"time"
)

type User struct {
	ID          string
	Name        string
	Password    string `json:"-"`
	Type        string `json:"-"`
	Face        string
	Parent      string
	Registetime interface{}
}

func (u *User) FixShow() *User {
	u.Registetime = (u.Registetime.(time.Time)).Format(utils.DateTimeFormart)
	return u
}
func DBPassword(password string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
