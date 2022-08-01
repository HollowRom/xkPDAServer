package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func getAllEmp(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetAllEmp(context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getEmp(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetEmp(context.Query(defNumberKey), context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
