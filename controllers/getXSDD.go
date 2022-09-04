package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func init() {
	AddHandlerGet("/getXSDDMain", getXSDDMain)
	AddHandlerGet("/getXSDDEntry", getXSDDEntry)
}

func getXSDDMain(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetXSDDMain(context.Query(defOrgKey), context.Query(defCustNumberKey), context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getXSDDEntry(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetXSDDEntry(context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
