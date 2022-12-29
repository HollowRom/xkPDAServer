package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"xkpdaserver/dbTools"
	"xkpdaserver/netTools"
)

func init() {
	AddHandlerPost("/postWorkShop", postWorkShop)
}

func postWorkShop(context *gin.Context) { // 定义请求接口和处理匿名函数
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

	miniStr := &dbTools.WorkShopGXJHMini{}
	e = json.Unmarshal(buf[0:i], miniStr)
	if e != nil {
		fmt.Println("解析post异常:" + e.Error())
		setErrJson(context, e)
		return
	}

	if miniStr.HeadMini == nil {
		fmt.Println("解析post异常,WorkShopHeadMini不能为空")
		setErrJson(context, e)
		return
	}

	if miniStr.EntryMini == nil {
		fmt.Println("解析post异常,WorkShopEntityMini不能为空")
		setErrJson(context, e)
		return
	}

	info := dbTools.GetPostWorkShop(miniStr)
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

	reb := netTools.PostPushSomeBill(infoJ)
	if reb == nil || len(reb) == 0 {
		setErrJson(context, nil)
		return
	}

	resp := &dbTools.ResponseStatus{}

	resp.SetStr(string(reb))

	//e = json.Unmarshal(reb, resp)
	//
	//if e != nil {
	//	setErrJson(context, e)
	//	return
	//}

	if !resp.IsSuccess() {
		setErrJson(context, errors.New(resp.GetErrMess()))
		return
	}

	gxhbList := dbTools.GetGXHBEntryMini(resp.GetReBillNo()[0], 0)
	if gxhbList == nil || len(gxhbList) == 0 {
		setErrJson(context, nil)
		return
	}

	var needInt []int
	var needFloat []float64
	for _, b := range miniStr.EntryMini {
		for _, b2 := range gxhbList {
			if b2.FOPERNUMBER == b.FOPERNUMBER {
				needInt = append(needInt, b2.FENTRYID)
				needFloat = append(needFloat, b.FQTY)
			}
		}
	}

	if len(needInt) == 0 {
		setErrJson(context, nil)
		return
	}

	var updMap = map[string]interface{}{}

	updMap["formid"] = "SFC_OperationReport"

	var entryMap []map[string]interface{}
	for i := range needInt {
		var tempMap = map[string]interface{}{}
		tempMap["FENTRYID"] = strconv.Itoa(needInt[i])
		tempMap["FFinishQty"] = strconv.FormatFloat(needFloat[i], 'f', 4, 64)
		tempMap["FQuaQty"] = strconv.FormatFloat(needFloat[i], 'f', 4, 64)
		entryMap = append(entryMap, tempMap)
	}

	var modelMap = map[string]interface{}{}

	modelMap["FID"] = gxhbList[0].FID
	modelMap["FEntity"] = entryMap

	var dataMap = map[string]interface{}{}

	dataMap["IsDeleteEntry"] = "true"
	dataMap["Model"] = modelMap

	updMap["data"] = dataMap

	reb, e = json.Marshal(updMap)

	if e != nil {
		setErrJson(context, e)
		return
	}

	reb = netTools.PostSaveSomeBill(reb)

	if reb == nil || len(reb) == 0 {
		setErrJson(context, nil)
		return
	}

	context.JSON(http.StatusOK, resp)
	netTools.AddCache(string(buf))
}
