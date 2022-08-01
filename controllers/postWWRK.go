package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"net/http"
	"strconv"
	"xkpdaserver/dbTools"
	"xkpdaserver/jsonTools"
	"xkpdaserver/netTools"
)

var (
	defWWRKFStockStatusId = "10000"
	//defWWRKFSrcBillTypeID  = "PUR_PurchaseOrder"
	defWWRKLinkFSTableName = "T_PUR_ReceiveEntry"
	defWWRKLinkFRuleId     = "PUR_ReceiveBill-STK_InStock"
)

func postWWRK(context *gin.Context) { // 定义请求接口和处理匿名函数
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
	miniStr := &jsonTools.WWRKMini{}
	e = json.Unmarshal(buf[0:i], miniStr)
	if e != nil {
		fmt.Println(e)
		setErrJson(context, e)
		return
	}

	if miniStr.WWRKHeadMini == nil {
		fmt.Println("解析post异常,QTCKHeadMini不能为空")
		setErrJson(context, e)
		return
	}

	if miniStr.WWRKEntityMini == nil {
		fmt.Println("解析post异常,QTCKEntityMini不能为空")
		setErrJson(context, e)
		return
	}

	for _, qm := range miniStr.WWRKEntityMini {
		if qm == nil {
			fmt.Println("解析post异常,QTCKEntityMini不能为空")
			setErrJson(context, e)
			return
		}
		if qm.FStockStatusId == "" {
			qm.FStockStatusId = defWWRKFStockStatusId
		}
		if len(qm.FLinkInfo) == 0 {
			qm.FLinkInfo = append(qm.FLinkInfo, map[string]string{})
		}
		if qm.FLinkInfo[0]["FInStockEntry_Link_FRuleId"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FRuleId"] = defWWRKLinkFRuleId
		}
		if qm.FLinkInfo[0]["FInStockEntry_Link_FSTableName"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FSTableName"] = defWWRKLinkFSTableName
		}

		if qm.FLinkInfo[0]["FInStockEntry_Link_FSBillId"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FSBillId"] = strconv.Itoa(qm.FPOORDERINTERID)
		}
		if qm.FLinkInfo[0]["FInStockEntry_Link_FSId"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FSId"] = strconv.Itoa(qm.FPOORDERENTRYID)
		}
		if qm.FLinkInfo[0]["FInStockEntry_Link_FBaseActualQty"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FBaseActualQty"] = qm.FQTY
		}
	}

	info := dbTools.GetPostWWRK(miniStr)
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

	reb := netTools.PostSome(defICUrl, infoJ)
	if reb == nil || len(reb) == 0 {
		setErrJson(context, nil)
		return
	}

	context.JSON(http.StatusOK, string(reb))
}
