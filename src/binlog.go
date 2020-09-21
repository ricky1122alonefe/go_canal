package main

import (
	"fmt"
	"github.com/siddontang/go-mysql/canal"
	"runtime/debug"
	"log"
	"github.com/ricky1122alonefe/go_canal/src/module"
)

type binlogHandler struct {
	canal.DummyEventHandler
	BinlogParser
}

func (h *binlogHandler) OnRow(e *canal.RowsEvent) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Print(r, " ", string(debug.Stack()))
		}
	}()



	// base value for canal.DeleteAction or canal.InsertAction
	var n = 0
	var k = 1

	if e.Action == canal.UpdateAction {
		n = 1
		k = 2
	}

	for i := n; i < len(e.Rows); i += k {
		fmt.Println(e.Rows[i])
		key := e.Table.Schema + "." + e.Table.Name

		switch key {
		case "PLATFORM.TB_RESOURCE":
			resource:=module.TBAResource{}
			h.GetBinLogData(&resource,e,i)
			switch e.Action {
			case canal.UpdateAction:
				oldUser := module.TBAResource{}
				h.GetBinLogData(&oldUser, e, i-1)
				log.Println(oldUser)
				fmt.Printf("name changed from %s to %s\n",oldUser,resource)
			case canal.InsertAction:

				fmt.Println("delete")
			case canal.DeleteAction:
				oldUser := module.TBAResource{}
				h.GetBinLogData(&oldUser, e, i-1)
				fmt.Printf("delete",oldUser.Id)
			}
		}

	}
	return nil
}

func (h *binlogHandler) String() string {
	return "binlogHandler"
}

func binlogListener() {
	c, err := getDefaultCanal()
	if err == nil {
		coords, err := c.GetMasterPos()
		if err == nil {
			c.SetEventHandler(&binlogHandler{})
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
