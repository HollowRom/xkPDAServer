package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func getAllKeeper(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetAllKeeper(context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getKeeper(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetKeeper(context.Query(defNumberKey), context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
