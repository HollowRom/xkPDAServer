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
	netTools.Init()
	dbTools.Init("driver={SQL Server};Server=127.0.0.1;Database=AIS20210805182552;user id=sa;password=sa;")
	db1 := dbTools.GetDB()

	defer func(db1 *xorm.Engine) {
		err := db1.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db1)

	if !netTools.TryLogin(nil) {
		fmt.Println("登录失败")
		panic(nil)
	}
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
