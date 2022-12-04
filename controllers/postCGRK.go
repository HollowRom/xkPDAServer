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
	defCGRKFStockStatusId  = "10000"
	//defCGRKFSrcBillTypeID  = "PUR_PurchaseOrder"
	defCGRKFSrcBillTypeID  = "PUR_PurchaseOrder"
	defCGRKLinkFSTableName = "t_PUR_POOrderEntry"
	//defCGRKLinkFRuleId     = "PUR_PurchaseOrder-STK_InStock"
	defCGRKLinkFRuleId     = "PUR_PurchaseOrder-STK_InStock"
)

func init() {
	AddHandlerPost("/postCGRK", postCGRK)
}

func postCGRK(context *gin.Context) { // 定义请求接口和处理匿名函数
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

	if netTools.GetCache(string(buf)) {
		fmt.Println("重复提交数据,一分钟后重试")
		setErrJson(context, e)
		return
	}

	miniStr := &dbTools.CGRKMini{}
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

	for _, qm := range miniStr.EntityMini {
		if qm == nil {
			fmt.Println("解析post异常,QTCKEntityMini不能为空")
			setErrJson(context, e)
			return
		}
		if qm.FStockStatusId == "" {
			qm.FStockStatusId = defCGRKFStockStatusId
		}
		if len(qm.FLinkInfo) == 0 {
			qm.FLinkInfo = append(qm.FLinkInfo, map[string]string{})
		}
		if qm.FSrcBillType == "" {
			qm.FSrcBillType = defCGRKFSrcBillTypeID
		}
		if qm.FLinkInfo[0]["FInStockEntry_Link_FRuleId"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FRuleId"] = defCGRKLinkFRuleId
		}
		if qm.FLinkInfo[0]["FInStockEntry_Link_FSTableName"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FSTableName"] = defCGRKLinkFSTableName
		}

		if qm.FLinkInfo[0]["FInStockEntry_Link_FSBillId"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FSBillId"] = strconv.Itoa(qm.FID)
			//qm.FLinkInfo[0]["FInStockEntry_Link_FSBillId"] = strconv.Itoa(qm.FSRCID)
		}
		if qm.FLinkInfo[0]["FInStockEntry_Link_FSId"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FSId"] = strconv.Itoa(qm.FENTRYID)
			//qm.FLinkInfo[0]["FInStockEntry_Link_FSId"] = strconv.Itoa(qm.FSRCENTRYID)
		}
		if qm.FLinkInfo[0]["FInStockEntry_Link_FBaseUnitQty"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FBaseUnitQty"] = qm.SQTY
		}
		if qm.FLinkInfo[0]["FInStockEntry_Link_FRemainInStockBaseQty"] == "" {
			qm.FLinkInfo[0]["FInStockEntry_Link_FRemainInStockBaseQty"] = qm.SQTY
		}
	}

	info := dbTools.GetPostCGRK(miniStr)
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
