package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func getAllXSDDMain(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetAllXSDDMain(context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
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
