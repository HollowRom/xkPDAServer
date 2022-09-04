package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func init() {
	AddHandlerGet("/getDepart", getDepart)
}

func getDepart(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetDepart(context.Query(defOrgKey), context.Query(defNumberKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
