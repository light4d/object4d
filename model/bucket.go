package model

import "time"

type bucket struct {
	Name string
	//R read,W write
	// '': belong to everyone;
	// 'groupname':belong to groupname;
	// 'username':belong to username;

	R, W       string
	Createtime time.Time
}
