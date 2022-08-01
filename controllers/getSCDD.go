package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func getAllSCDDMain(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetAllSCDDMain(context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getSCDDMain(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetSCDDMain(context.Query(defOrgKey), context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getSCDDEntry(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetSCDDEntry(context.Query(defBillKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
