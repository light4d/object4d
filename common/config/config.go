package config

import (
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Object4dPort int    `json:"object4d_port"`
	Mysql        string `json:"mysql"`
	Redis        struct {
		Addr     string `json:"addr"`
		Password string `json:"password"`
		DB       int    `json:"db"`
	} `json:"redis"`
}
