package model

import "time"

type Object struct {
	Name   string
	Bucket string
	Minio  int
	//R read,W write
	// '': belong to everyone;
	// 'groupname':belong to groupname;
	R, W       string
	Createtime time.Time
}
