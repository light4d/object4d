package model

import "time"

type Object struct {
	Name string
	//R read,W write
	// '': belong to everyone;
	// 'groupname':belong to groupname;
	R, W       string
	Createtime time.Time
}
