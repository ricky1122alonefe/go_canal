package main

import (
	"log"

	"github.com/ricky1122alonefe/go_canal/src/conf"
)

func main() {
	c, err := getDefaultCanal()
	if err != nil {
		log.Println(err.Error())
	}
	var config conf.CananConfig
	schemaMap := make(map[string][]string)
	list := []string{"1", "2"}
	config.Address = ""
	schemaMap["1"] = list
	config.SchemaInfo = schemaMap

	var binLog BinlogHandler
	binLog.Config = config
	binLog.InitSchema()
	c.SetEventHandler(&BinlogHandler{Config: config})
	pos, err1 := c.GetMasterPos()
	if err1 != nil {
		log.Println(err1.Error())
	}
	c.RunFrom(pos)
}
