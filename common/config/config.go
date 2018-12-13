package config

import (
	_ "github.com/go-sql-driver/mysql"
)

type Config struct {
	Object4dPort int    `json:"object4d_port"`
	Mysql        string `json:"mysql"`
}
