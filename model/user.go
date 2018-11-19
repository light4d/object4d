package model

import "time"

type User struct {
	ID          string
	Name        string
	Password    string
	Face        string
	RegisteTime time.Time
}

func (u *User) FixShow() *User {
	return u
}
