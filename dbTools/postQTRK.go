package dbTools

import (
	"fmt"
	"time"
)

const (
	defQTRKBillType = "d772ead981e748d69dda1caac7583f8c"
	defQTRKFromId   = "STK_MISCELLANEOUS"
	defQTRKFStockStatusID = "10000"
)

func GetPostQTRK(mini *QTRKMini) ModelBaseInterface {
	if mini == nil || mini.EntityMini == nil || len(mini.EntityMini) < 0 {
		fmt.Println("qtrk入参异常")
		return nil
	}

	for idx := 0; idx < len(mini.EntityMini); idx++ {
		if mini.EntityMini[idx].FStockStatusId == "" {
			mini.EntityMini[idx].FStockStatusId = defQTRKFStockStatusID
		}
	}

	i := InitQTRKModel(&DefModelHeadBase{FBillTypeId: defQTRKBillType, FDate: time.Now(), FromId: defQTRKFromId})

	if i == nil {
		return nil
	}

	i.AddModelHead(mini.EntityMini)

	i.AddModelFEntities(mini.EntityMini)
	return i
}
