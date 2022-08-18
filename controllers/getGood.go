package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func getAllGood(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetAllGood(context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getGood(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetGood(context.Query(defOrgKey), context.Query(defNumberKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
