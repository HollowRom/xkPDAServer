package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func getAllWWTLMain(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetAllWWTLMain(context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getWWTLMain(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetWWTLMain(context.Query(defOrgKey), context.Query(defSuppKey), context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getWWTLEntry(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetWWTLEntry(context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
