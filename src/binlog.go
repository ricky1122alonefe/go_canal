package main

import (
	"fmt"
	"runtime/debug"

	"github.com/ricky1122alonefe/go_canal/src/conf"
	"github.com/rickyalonefe1122/go_canal/src/module"
	"github.com/siddontang/go-mysql/canal"
)

var schemaList []string

func (h *BinlogHandler) InitSchema() {
	for k, v := range h.Config.SchemaInfo {
		for _, table := range v {
			str := k + "." + table
			schemaList = append(schemaList, str)
		}
	}
}

type BinlogHandler struct {
	canal.DummyEventHandler
	BinlogParser
	Config conf.CananConfig
}

func (h *BinlogHandler) OnRow(e *canal.RowsEvent) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r, " ", string(debug.Stack()))
		}
	}()

	// base value for canal.DeleteAction or canal.InsertAction
	var n = 0
	var k = 1

// 	if e.Action == canal.UpdateAction {
// 		n = 1
// 		k = 2
// 	}
// 		switch key {
// 		// case "PLATFORM.TB_RESOURCE":
// 		// 	resource:=module.TBAResource{}
// 		// 	h.GetBinLogData(&resource,e,i)
// 		// 	switch e.Action {
// 		// 	case canal.UpdateAction:
// 		// 		oldResource := module.TBAResource{}
// 		// 		h.GetBinLogData(&oldResource, e, i-1)
// 		// 		log.Println(oldResource)
// 		// 		fmt.Printf("name changed from %s to %s\n",oldResource,resource)
// 		// 	case canal.InsertAction:
// 		//
// 		// 		fmt.Println("delete")
// 		// 	case canal.DeleteAction:  11111
// 		// 		oldUser := module.TBAResource{}
// 		// 		h.GetBinLogData(&oldUser, e, i-1)
// 		// 		fmt.Printf("delete",oldUser.Id)
// 		// 	}
// 		}
// 	}
	return nil
}

func (h *BinlogHandler) String() string {
	return "binlogHandler"
}

func binlogListener() {
	c, err := getDefaultCanal()
	if err == nil {
		coords, err := c.GetMasterPos()
		if err == nil {
			c.SetEventHandler(&BinlogHandler{})
			c.RunFrom(coords)
		}
	}
}

func getDefaultCanal() (*canal.Canal, error) {
	cfg := canal.NewDefaultConfig()
	// cfg.Addr = fmt.Sprintf("%s:%d", "mariadb", 3307)
	// cfg.User = "root"
	// cfg.Password = "root"
	// cfg.Flavor = "mysql"
	cfg.Addr = "127.0.0.1:3307"
	cfg.User = "root"
	cfg.Password = "145900"

	cfg.Dump.ExecutionPath = ""

	return canal.NewCanal(cfg)
}
