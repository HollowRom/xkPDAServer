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
	defSCRKFStockStatusId  = "10000"
	defSCRKFSrcBillTypeID  = "PUR_ReceiveBill"
	defSCRKLinkFSTableName = "T_PRD_MOENTRY"
	defSCRKLinkFRuleId     = "PRD_MO2INSTOCK"
)

func init() {
	AddHandlerPost("/postSCRK", postSCRK)
}

func postSCRK(context *gin.Context) { // 定义请求接口和处理匿名函数
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
	miniStr := &dbTools.SCRKMini{}
	e = json.Unmarshal(buf[0:i], miniStr)
	if e != nil {
		fmt.Println("解析post异常:" + e.Error())
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
			qm.FStockStatusId = defSCRKFStockStatusId
		}
		//if qm.FBILLNO != "" && qm.FLinkInfo != nil {
			if len(qm.FLinkInfo) == 0 {
				qm.FLinkInfo = append(qm.FLinkInfo, map[string]string{})
			}
			if qm.FSrcBillType == "" {
				qm.FSrcBillType = defSCRKFSrcBillTypeID
			}
			if qm.FLinkInfo[0]["FENTITY_Link_FRuleId"] == "" {
				qm.FLinkInfo[0]["FENTITY_Link_FRuleId"] = defSCRKLinkFRuleId
			}
			if qm.FLinkInfo[0]["FENTITY_Link_FSTableName"] == "" {
				qm.FLinkInfo[0]["FENTITY_Link_FSTableName"] = defSCRKLinkFSTableName
			}
			if qm.FLinkInfo[0]["FENTITY_Link_FSBillId"] == "" {
				qm.FLinkInfo[0]["FENTITY_Link_FSBillId"] = strconv.Itoa(qm.FID)
			}
			if qm.FLinkInfo[0]["FENTITY_Link_FSId"] == "" {
				qm.FLinkInfo[0]["FENTITY_Link_FSId"] = strconv.Itoa(qm.FENTRYID)
			}
			if qm.FLinkInfo[0]["FENTITY_Link_FBaseActualQty"] == "" {
				qm.FLinkInfo[0]["FENTITY_Link_FBaseActualQty"] = qm.SQTY
			}
		//}
	}
	//fmt.Println("miniStr.HeadMini:", *miniStr.HeadMini)
	//
	//for idx := 0; idx < len(miniStr.EntityMini) - 1; idx++ {
	//	fmt.Println(idx, ":miniStr.EntityMini:", *miniStr.EntityMini[idx])
	//}

	info := dbTools.GetPostSCRK(miniStr)
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
