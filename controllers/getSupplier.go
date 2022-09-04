package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"xkpdaserver/dbTools"
)

func init() {
	AddHandlerGet("/getSupplier", getSupplier)
}

func getSupplier(context *gin.Context) { // 定义请求接口和处理匿名函数
	info := dbTools.GetSupplier(context.Query(defOrgKey), context.Query(defNumberKey))
	if info == nil {
		fmt.Println("supp == nil")
		setErrJson(context, nil)
		return
	}
	context.JSON(http.StatusOK, info)
}
