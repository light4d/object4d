package model

import "time"

type User struct {
	ID         string
	Name       string
	Password   string
	Face       string
	Createtime time.Time
}
