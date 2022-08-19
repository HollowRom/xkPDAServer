package dbTools

import (
	"fmt"
	"time"
)

const (
	defQTRKBillType = "d772ead981e748d69dda1caac7583f8c"
	defQTRKFromId   = "STK_MISCELLANEOUS"
)

func GetPostQTRK(mini *QTRKMini) ModelBaseInterface {
	if mini == nil {
		fmt.Println("qtrk入参异常")
		return nil
	}
	i := InitQTRKModel(&DefModelHeadBase{FBillTypeId: defQTRKBillType, FDate: time.Now(), FromId: defQTRKFromId})

	if i == nil {
		return nil
	}

	//i.AddModelHead(mini.QTRKHeadMini)

	i.AddModelFEntities(mini.QTRKEntityMini, mini.QTRKEntityMini[0].FUseOrgNumber)
	return i
}
