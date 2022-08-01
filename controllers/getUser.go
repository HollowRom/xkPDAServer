package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func getAllUser(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetAllUser()
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}

func getUser(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetUser(context.Query(defNumberKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
