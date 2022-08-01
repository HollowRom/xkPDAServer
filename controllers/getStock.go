package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func getAllStock(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetAllStock(context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getStock(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetStock(context.Query(defNumberKey), context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
