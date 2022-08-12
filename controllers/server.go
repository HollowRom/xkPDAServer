package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

var engine = gin.Default()

const (
	defICUrl       = "http://127.0.0.1/k3cloud/Kingdee.BOS.WebApi.ServicesStub.DynamicFormService.Save.common.kdsvc"
	defReadBufSize = 10240
	defOrgKey      = "FUseOrgNumber"
	defBillKey     = "FBILLNO"
	defSuppKey     = "FSupplierNumber"
	defNumberKey   = "FNUMBER"
)

var defPort = ":8090"

var returnEmptyJSOn = &gin.H{}

func setErrJson(c *gin.Context, e error) {
	if e != nil {
		c.JSON(http.StatusBadRequest, &gin.H{"err": e.Error()})
	} else {
		c.JSON(http.StatusBadRequest, returnEmptyJSOn)
	}
}

func Init() {
	tempValue := dbTools.GetConfFromKey("listenPort")

	if tempValue[0] != ':' {
		defPort = ":" + tempValue
	} else {
		defPort = tempValue
	}

	engine.GET("/ping", ping)

	engine.GET("/getGood", getGood)

	engine.GET("/getAllGood", getAllGood)

	engine.GET("/getCustomer", getCustomer)

	engine.GET("/getAllCustomer", getAllCustomer)

	engine.GET("/getEmp", getEmp)

	engine.GET("/getAllEmp", getAllEmp)

	engine.GET("/getStock", getStock)

	engine.GET("/getAllStock", getAllStock)

	engine.GET("/getSupplier", getSupplier)

	engine.GET("/getAllSupplier", getAllSupplier)

	engine.GET("/getUser", getUser)

	engine.GET("/getAllKeeper", getAllKeeper)

	engine.GET("/getKeeper", getKeeper)

	engine.GET("/getAllOrg", getAllOrg)

	engine.GET("/getOrg", getOrg)

	engine.GET("/getAllCGDDMain", getAllCGDDMain)

	engine.GET("/getCGDDMain", getCGDDMain)

	engine.GET("/getCGDDEntry", getCGDDEntry)

	engine.GET("/getAllSCDDMain", getAllSCDDMain)

	engine.GET("/getSCDDMain", getSCDDMain)

	engine.GET("/getSCDDEntry", getSCDDEntry)

	engine.GET("/getAllXSDDMain", getAllXSDDMain)

	engine.GET("/getXSDDMain", getXSDDMain)

	engine.GET("/getXSDDEntry", getXSDDEntry)

	engine.GET("/getAllSCTLMain", getAllSCTLMain)

	engine.GET("/getSCTLMain", getSCTLMain)

	engine.GET("/getSCTLEntry", getSCTLEntry)

	engine.GET("/getAllWWTLMain", getAllWWTLMain)

	engine.GET("/getWWTLMain", getWWTLMain)

	engine.GET("/getWWTLEntry", getWWTLEntry)

	engine.GET("/getAllWWDDMain", getAllWWDDMain)

	engine.GET("/getWWDDMain", getWWDDMain)

	engine.GET("/getWWDDEntry", getWWDDEntry)

	engine.GET("/getAllUser", getAllUser)

	engine.POST("/postPing", postPing)

	engine.POST("/postQTCK", postQTCK)

	engine.POST("/postQTRK", postQTRK)

	engine.POST("/postCGRK", postCGRK)

	engine.POST("/postSCRK", postSCRK)

	engine.POST("/postSCLL", postSCLL)

	engine.POST("/postXSCK", postXSCK)

	engine.POST("/postWWLL", postWWLL)

	engine.POST("/postWWRK", postWWRK)

	fmt.Println("注册路由完成尝试启动监听端口" + defPort)

	if err := engine.Run(defPort); err != nil {
		panic(err)
	}
}

func ping(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "pong",
	})
}

func postPing(context *gin.Context) {
	context.JSON(http.StatusOK, gin.H{
		"message": "postPong",
	})
}
