package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func init() {
	AddHandlerGet("/getCGDDMain", getCGDDMain)
	AddHandlerGet("/getCGDDEntry", getCGDDEntry)
}
func getCGDDMain(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetCGDDMain(context.Query(defOrgKey), context.Query(defSuppKey), context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getCGDDEntry(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetCGDDEntry(context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
