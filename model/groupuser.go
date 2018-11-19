package model

import "time"

type GroupUser struct {
	ID       string
	User     string
	Jointime time.Time
}
