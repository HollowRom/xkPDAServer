package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func getAllSupplier(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetAllSupplier()
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getSupplier(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetSupplier(context.Query(defNumberKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
