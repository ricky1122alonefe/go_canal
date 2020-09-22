package main

import "log"

func main() {
	c,err:=getDefaultCanal()
	if err !=nil{
		log.Println(err.Error())
	}



	c.SetEventHandler(&binlogHandler{})
	pos,err1:=c.GetMasterPos()
	if err1!=nil{
		log.Println(err1.Error())
	}
	c.RunFrom(pos)
}
