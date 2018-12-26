package server

import (
	"encoding/json"
	"github.com/gobestsdk/gobase/log"
	"github.com/light4d/object4d/common/config"
	"io/ioutil"
)

var (
	APPConfig config.Config = config.Config{
		Mysql:        "root:@tcp(localhost:3306)/object4d?charset=utf8mb4&parseTime=true",
		Object4dPort: 9001,
	}
)

func ParseConfig(configfilepath string) error {
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
