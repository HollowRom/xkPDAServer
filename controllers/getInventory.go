package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func init() {
	AddHandlerGet("/getInventory", getInventory)
}

func getInventory(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetInventory(context.Query(defOrgKey), context.Query(defNumberKey), context.Query(defLostTextKey), context.Query(defStockNumberKey))
	if info == nil {
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}