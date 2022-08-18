package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func getAllCustomer(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetAllCustomer(context.Query(defOrgKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getCustomer(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetCustomer(context.Query(defOrgKey), context.Query(defNumberKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
