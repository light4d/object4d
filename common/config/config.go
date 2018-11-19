package config

import (
	"encoding/json"
	_ "github.com/go-sql-driver/mysql"
	"github.com/gobestsdk/gobase/log"
	"io/ioutil"
)

type Config struct {
	HttpPort int    `json:"http_port"`
	Mysql    string `json:"mysql"`
	Weixin   struct {
		APPID  string `json:"appid"`
		SECRET string `json:"secret"`
	}
	Qiniu struct {
		Ak string `json:"ak"`
		Sk string `json:"sk"`
	} `json:"qiniu"`
}

var (
	APPConfig Config = Config{
		HttpPort: 8000,
		Mysql:    "",
	}
)

func Init(configfilepath string) error {
	data, err := ioutil.ReadFile(configfilepath)
	if err != nil {
		log.Fatal(log.Fields{"error": err, "app": "config file read "})
		return err
	}

	err = json.Unmarshal([]byte(data), &APPConfig)
	if err != nil {
		log.Fatal(log.Fields{"error": err, "app": "config file parse "})
		return err
	}
	log.Info(log.Fields{"app": "config file", "config": APPConfig})
	return nil
}
