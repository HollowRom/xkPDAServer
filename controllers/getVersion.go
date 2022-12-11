package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"sync"
)

var serverVersion = "version.100002"

var vl = &sync.RWMutex{}

func init() {
	AddHandlerGet("/getVersion", getVersion)
}

func gv() string {
	vl.RLock()
	defer vl.RUnlock()
	return serverVersion
}

func getVersion(context *gin.Context) { // 定义请求接口和处理匿名函数
	context.JSON(http.StatusOK, gv())
}
