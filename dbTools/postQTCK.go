package dbTools

import (
	"fmt"
	"time"
	"xkpdaserver/jsonTools"
)

const (
	defQTCKBillType = "54533291F9A44D38809F70000499BEE9"
	defQTCKFromId   = "STK_MisDelivery"
)

func GetPostQTCK(mini *jsonTools.QTCKMini) jsonTools.ModelBaseInterface {
	if mini == nil || mini.QTCKHeadMini == nil || mini.QTCKHeadMini.FOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := jsonTools.InitQTCKModel(&jsonTools.DefModelHeadBase{FBillTypeId: defQTCKBillType, FDate: time.Now(), FromId: defQTCKFromId})

	if i == nil {
		return nil
	}

	i.AddModelHead(mini.QTCKHeadMini)

	i.AddModelFEntities(mini.QTCKEntityMini, mini.QTCKHeadMini.FOrgNumber)
	return i
}
