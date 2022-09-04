package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func init() {
	AddHandlerGet("/getStock", getStock)
}

func getStock(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetStock(context.Query(defOrgKey), context.Query(defNumberKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
