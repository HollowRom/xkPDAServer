package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func init() {
	AddHandlerGet("/getOrg", getOrg)
}

func getOrg(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetOrg(context.Query(defNumberKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
