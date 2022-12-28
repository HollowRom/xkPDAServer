package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"xkpdaserver/dbTools"
	"xkpdaserver/netTools"
)

var (
	defWWLLFStockStatusId = "10000"
	//defWWLLFSrcBillTypeID  = "SUB_PPBOM"
	defWWLLLinkFSTableName = "T_SUB_PPBOMENTRY"
	defWWLLLinkFRuleId     = "SUB_PPBOM_Pick"
)

func init() {
	AddHandlerPost("/postWWLL", postWWLL)
}

func postWWLL(context *gin.Context) { // 定义请求接口和处理匿名函数
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
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(context.Request.Body)
	if netTools.GetCacheWithCTX(string(buf), context) {
		return
	}

	miniStr := &dbTools.WWLLMini{}
	e = json.Unmarshal(buf[0:i], miniStr)
	if e != nil {
		fmt.Println(e)
		setErrJson(context, e)
		return
	}

	if miniStr.HeadMini == nil {
		fmt.Println("解析post异常,QTCKHeadMini不能为空")
		setErrJson(context, e)
		return
	}

	if miniStr.EntityMini == nil {
		fmt.Println("解析post异常,QTCKEntityMini不能为空")
		setErrJson(context, e)
		return
	}

	for _, qm := range miniStr.EntityMini {
		if qm == nil {
			fmt.Println("解析post异常,QTCKEntityMini不能为空")
			setErrJson(context, e)
			return
		}
		if qm.FStockStatusId == "" {
			qm.FStockStatusId = defWWLLFStockStatusId
		}
		if len(qm.FLinkInfo) == 0 {
			qm.FLinkInfo = append(qm.FLinkInfo, map[string]string{})
		}
		if qm.FLinkInfo[0]["FEntity_Link_FRuleId"] == "" {
			qm.FLinkInfo[0]["FEntity_Link_FRuleId"] = defWWLLLinkFRuleId
		}
		if qm.FLinkInfo[0]["FEntity_Link_FSTableName"] == "" {
			qm.FLinkInfo[0]["FEntity_Link_FSTableName"] = defWWLLLinkFSTableName
		}

		if qm.FLinkInfo[0]["Entity_Link_FSBillId"] == "" {
			qm.FLinkInfo[0]["Entity_Link_FSBillId"] = strconv.Itoa(qm.FSrcInterId)
		}
		if qm.FLinkInfo[0]["FEntity_Link_FSId"] == "" {
			qm.FLinkInfo[0]["FEntity_Link_FSId"] = strconv.Itoa(qm.FSrcEntryId)
		}
		if qm.FLinkInfo[0]["FEntity_Link_FBaseActualQty"] == "" {
			qm.FLinkInfo[0]["FEntity_Link_FBaseActualQty"] = qm.FMustQty
		}
	}

	info := dbTools.GetPostWWLL(miniStr)
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
