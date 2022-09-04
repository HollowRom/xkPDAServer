package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func init() {
	AddHandlerGet("/getSCTLMain", getSCTLMain)
	AddHandlerGet("/getSCTLEntry", getSCTLEntry)
}

func getSCTLMain(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetSCTLMain(context.Query(defOrgKey), context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getSCTLEntry(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetSCTLEntry(context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
