package model

import "time"

type Group struct {
	ID   string
	Name string
	//Owner 管理员ID
	Owner      string
	Password   string
	Face       string
	Createtime time.Time
}
