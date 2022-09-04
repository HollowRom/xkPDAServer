package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func init() {
	AddHandlerGet("/getWWDDMain", getWWDDMain)
	AddHandlerGet("/getWWDDEntry", getWWDDEntry)
}

func getWWDDMain(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetWWDDMain(context.Query(defOrgKey), context.Query(defSuppKey), context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getWWDDEntry(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetWWDDEntry(context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
