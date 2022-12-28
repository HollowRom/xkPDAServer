package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"xkpdaserver/dbTools"
	"xkpdaserver/netTools"
)

var (
	defQTRKFStockDirect   = "GENERAL"
	defQTRKFStockStatusId = "10000"
)

func init() {
	AddHandlerPost("/postQTRK", postQTRK)
}

func postQTRK(context *gin.Context) { // 定义请求接口和处理匿名函数
	buf := make([]byte, defReadBufSize)
	i, e := context.Request.Body.Read(buf)
	if e != nil && e != io.EOF {
		fmt.Println("读取post异常:" + e.Error())
		setErrJson(context, e)
		return
	}
	if buf[defReadBufSize-1] != 0 {
		buf = make([]byte, defReadBufSize*2)
		i, e = context.Request.Body.Read(buf)
		if e != nil && e != io.EOF {
			fmt.Println("读取post异常:" + e.Error())
			setErrJson(context, e)
			return
		}
	}
	fmt.Println("接受到的post信息:" + string(buf[0:i]))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(context.Request.Body)

	if netTools.GetCacheWithCTX(string(buf), context) {
		return
	}

	miniStr := &dbTools.QTRKMini{}
	e = json.Unmarshal(buf[0:i], miniStr)
	if e != nil {
		fmt.Println(e)
		setErrJson(context, e)
		return
	}

	//if miniStr.HeadMini == nil {
	//	fmt.Println("解析post异常,QTCKHeadMini不能为空")
	//	setErrJson(context, e)
	//	return
	//}

	if miniStr.EntityMini == nil {
		fmt.Println("解析post异常,QTCKEntityMini不能为空")
		setErrJson(context, e)
		return
	}

	//if miniStr.HeadMini.FStockDirect == "" {
	//	miniStr.HeadMini.FStockDirect = defQTRKFStockDirect
	//}

	for _, qm := range miniStr.EntityMini {
		if qm == nil {
			fmt.Println("解析post异常,QTCKEntityMini不能为空")
			setErrJson(context, e)
			return
		}
		if qm.FStockStatusId == "" {
			qm.FStockStatusId = defQTRKFStockStatusId
		}
	}

	info := dbTools.GetPostQTRK(miniStr)
	if info == nil {
		fmt.Println("mini生成完整json返回为空")
		setErrJson(context, nil)
		return
	}
	infoJ := info.GetJson()
	if infoJ == nil {
		setErrJson(context, nil)
		return
	}

	reb := netTools.PostSaveSomeBill(infoJ)
	if reb == nil || len(reb) == 0 {
		setErrJson(context, nil)
		return
	}

	resp := &dbTools.ResponseStatus{}

	e = json.Unmarshal(reb, resp)

	if e != nil {
		setErrJson(context, nil)
		return
	}

	context.JSON(http.StatusOK, resp)
	netTools.AddCache(string(buf))
}
