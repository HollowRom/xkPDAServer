package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func init() {
	AddHandlerGet("/loadingNewApp", loadingNewApp)
}
func loadingNewApp(context *gin.Context) { // 定义请求接口和处理匿名函数
	filename := "./files/app-release.apk"
	context.Writer.Header().Add("Content-Disposition", fmt.Sprintf("attachment; filename=%s", filename))
	context.Writer.Header().Add("Content-Type", "application/octet-stream")
	context.File(filename)
}
