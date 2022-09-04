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
	defXSCKFStockStatusId  = "10000"
	defXSCKLinkFSTableName = "T_SAL_ORDERENTRY"
	defXSCKLinkFRuleId     = "SaleOrder-OutStock"
)

func init() {
	AddHandlerPost("/postXSCK", postXSCK)
}

func postXSCK(context *gin.Context) { // 定义请求接口和处理匿名函数
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
	fmt.Println("收到的post:" + string(buf[0:i]))
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
		}
	}(context.Request.Body)
	miniStr := &dbTools.XSCKMini{}
	e = json.Unmarshal(buf[0:i], miniStr)
	if e != nil {
		fmt.Println(e)
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
			qm.FStockStatusId = defXSCKFStockStatusId
		}
		if len(qm.FLinkInfo) == 0 {
			qm.FLinkInfo = append(qm.FLinkInfo, map[string]string{})
		}
		if qm.FLinkInfo[0]["FEntity_Link_FRuleId"] == "" {
			qm.FLinkInfo[0]["FEntity_Link_FRuleId"] = defXSCKLinkFRuleId
		}
		if qm.FLinkInfo[0]["FEntity_Link_FSTableName"] == "" {
			qm.FLinkInfo[0]["FEntity_Link_FSTableName"] = defXSCKLinkFSTableName
		}
		if qm.FLinkInfo[0]["FEntity_Link_FSBillId"] == "" {
			qm.FLinkInfo[0]["FEntity_Link_FSBillId"] = strconv.Itoa(qm.FID)
		}
		if qm.FLinkInfo[0]["FEntity_Link_FSId"] == "" {
			qm.FLinkInfo[0]["FEntity_Link_FSId"] = strconv.Itoa(qm.FENTRYID)
		}
		if qm.FLinkInfo[0]["FEntity_Link_FBaseUnitQty"] == "" {
			qm.FLinkInfo[0]["FEntity_Link_FBaseUnitQty"] = qm.SQTY
		}
		if qm.FLinkInfo[0]["FEntity_Link_FSALBASEQTY"] == "" {
			qm.FLinkInfo[0]["FEntity_Link_FSALBASEQTY"] = qm.SQTY
		}
	}

	info := dbTools.GetPostXSCK(miniStr)
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

	resp := &dbTools.ResponseStatus{}

	e = json.Unmarshal(reb, resp)

	if e != nil {
		setErrJson(context, nil)
		return
	}

	context.JSON(http.StatusOK, resp)
}
