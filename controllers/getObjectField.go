package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
	"time"
	"xkpdaserver/netTools"
)

const objectFieldVersionKey = "version"
const checkVersionJsonChangeTime = time.Second * 60

var objectFieldJson *netTools.ObjectFieldData

var ofl = &sync.RWMutex{}

func updVersionFieldJson() {
	v, j := netTools.ReadVersionFieldJson()
	ov := gv()
	if ov != "" && v != gv() {
		ofl.Lock()
		defer ofl.Unlock()
		vl.Lock()
		defer vl.Unlock()
		serverVersion = v
		objectFieldJson = j
	}
}

func getOJ() *netTools.ObjectFieldData {
	ofl.RUnlock()
	defer ofl.RUnlock()
	return objectFieldJson
}

func init() {
	AddHandlerGet("/getObjectField", getObjectField)
	go func() {
		for {
			updVersionFieldJson()
			time.Sleep(checkVersionJsonChangeTime)
		}
	}()
}

func getObjectField(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := context.Query(objectFieldVersionKey)
	if info == "" {
		setErrJson(context, nil)
		return
	}
	if info == gv() {
		context.JSON(http.StatusOK, getOJ())
		return
	}
	setErrJson(context, nil)
}
