package model

import (
	"crypto/md5"
	"encoding/hex"
	"time"
)

type User struct {
	ID          string
	Name        string
	Password    string `json:"-"`
	Face        string
	RegisteTime time.Time
}

func (u *User) FixShow() *User {

	return u
}
func DBPassword(password string) string {
	md5Ctx := md5.New()
	md5Ctx.Write([]byte(password))
	cipherStr := md5Ctx.Sum(nil)
	return hex.EncodeToString(cipherStr)
}
