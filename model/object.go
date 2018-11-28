package model

import "time"

type Object struct {
	Object4d
	Minio  int
	Name   string
	Folder string
	//R read,W write
	// '': belong to everyone;
	// 'groupname':belong to groupname;
	R, W       string
	Createtime time.Time
}
