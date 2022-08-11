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
	//netTools.Init()
	dbTools.Init("driver={SQL Server};Server=127.0.0.1;Database=AIS20210805182552;user id=sa;password=sa;")
	db1 := dbTools.GetDB()

	defer func(db1 *xorm.Engine) {
		err := db1.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(db1)

	//db := dbTools.GetDB()

	var cgMain []*dbTools.GoodsInto

	//ss := db1.Where(fmt.Sprintf("FUseOrgNumber = '%d'", 100)) //.Find(&cgMain)
	//_ = ss.Find(&cgMain)
	cgMain = dbTools.GetGood("3.03", "100")

	for idx := 0; idx < len(cgMain); idx++ {
		fmt.Println(*cgMain[idx])
	}
	fmt.Println("cgmain:", len(cgMain))
	return
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
