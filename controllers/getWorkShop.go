package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func init() {
	AddHandlerGet("/getWorkShopMain", getWorkShopMain)

	AddHandlerGet("/getWorkShopEntry", getWorkShopEntry)
}

func getWorkShopMain(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetWorkShopMain(context.Query(defOrgKey), context.Query(defBillKey), context.Query(defNumberKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getWorkShopEntry(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetWorkShopEntry(context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
