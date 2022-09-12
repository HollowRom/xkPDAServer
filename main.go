package main

import (
	"fmt"
	"strconv"
	"time"
	"xkpdaserver/controllers"
	"xkpdaserver/dbTools"
	"xkpdaserver/netTools"
	"xorm.io/xorm"
)

func main() {
	someInit()
}

func someInit() {
	//涉及读取配置文件生成url所以需要优先初始化
	netTools.Init()
	dbTools.Init("driver={SQL Server};Server=127.0.0.1;Database=AIS20210805182552;user id=sa;password=sa;")

	defer func(d *xorm.Engine) {
		e := d.Close()
		if e != nil {
			fmt.Println(e)
		}
	}(dbTools.GetDB())

	tempValue := dbTools.GetConfFromKey("serverTimeOut")
	if tempValue != "" && tempValue != "-1" {
		timeOut, _ := strconv.Atoi(tempValue)
		go func() {
			time.Sleep(time.Duration(timeOut) * time.Second)
			panic(nil)
		}()
	}

	controllers.Init()
}
