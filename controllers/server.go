package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"xkpdaserver/dbTools"
)

const (
	defReadBufSize    = 10240
	defOrgKey         = "FOrgNumber"
	defBillKey        = "FBillNo"
	defSuppKey        = "FSupplierNumber"
	defNumberKey      = "FNumber"
	defGoodNumberKey  = "FGoodNumber"
	defStockNumberKey = "FStockNumber"
	defCustNumberKey  = "FCustNumber"
	defLostTextKey    = "FLotText"
	defHost = "http://121.37.169.235"
)

var engine = gin.Default()

var defPort = ":8090"

var returnEmptyJSOn = &gin.H{}

var handlerGetMap = map[string]func(*gin.Context){}

var handlerPostMap = map[string]func(*gin.Context){}

func setErrJson(c *gin.Context, e error) {
	if e != nil {
		c.JSON(http.StatusBadRequest, &gin.H{"err": e.Error()})
	} else {
		c.JSON(http.StatusBadRequest, returnEmptyJSOn)
	}
}

func AddHandlerGet(k string, v func(*gin.Context)) {
	handlerGetMap[k] = v
}

func AddHandlerPost(k string, v func(*gin.Context)) {
	handlerPostMap[k] = v
}

var o sync.Once

func init() {
	AddHandlerGet("/ping", ping)
	AddHandlerPost("/postPing", postPing)
}

var onceInit = func (){
	tempValue := dbTools.GetConfFromKey("listenPort")

	if tempValue[0] != ':' {
		defPort = ":" + tempValue
	} else {
		defPort = tempValue
	}

	for k, v := range handlerGetMap {
		engine.GET(k, v)
	}

	for k, v := range handlerPostMap {
		engine.POST(k, v)
	}

	fmt.Println("注册路由完成尝试启动监听端口" + defPort)

	if err := engine.Run(defPort); err != nil {
		panic(err)
	}
}

func Init() {
	o.Do(onceInit)
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
