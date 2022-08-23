package dbTools

import (
	"fmt"
	"time"
)

const (
	defQTCKBillType = "54533291F9A44D38809F70000499BEE9"
	defQTCKFromId   = "STK_MisDelivery"
)

func GetPostQTCK(mini *QTCKMini) ModelBaseInterface {
	if mini == nil || mini.EntityMini == nil || len(mini.EntityMini) < 0 {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := InitQTCKModel(&DefModelHeadBase{FBillTypeId: defQTCKBillType, FDate: time.Now(), FromId: defQTCKFromId})

	if i == nil {
		return nil
	}

	i.AddModelHead(mini.EntityMini[0])

	i.AddModelFEntities(mini.EntityMini)
	return i
}
