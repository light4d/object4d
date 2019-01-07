package dao

import (
	"github.com/jinzhu/gorm"
)

func DB(Mysql string) *gorm.DB {
	//db, err := gorm.Open("mysql", "user:password@/dbname?charset=utf8&parseTime=True&loc=Local")
	db, err := gorm.Open("mysql", Mysql)
	db.LogMode(true)

	if err != nil {
		panic(err)
	}
	defer db.Close()
	return db
}
