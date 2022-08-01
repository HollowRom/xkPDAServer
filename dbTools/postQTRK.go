package dbTools

import (
	"fmt"
	"time"
	"xkpdaserver/jsonTools"
)

const (
	defQTRKBillType = "d772ead981e748d69dda1caac7583f8c"
	defQTRKFromId   = "STK_MISCELLANEOUS"
)

func GetPostQTRK(mini *jsonTools.QTRKMini) jsonTools.ModelBaseInterface {
	if mini == nil || mini.QTRKHeadMini == nil || mini.QTRKHeadMini.FOrgNumber == "" {
		fmt.Println("qtrk入参异常")
		return nil
	}
	i := jsonTools.InitQTRKModel(&jsonTools.DefModelHeadBase{FBillTypeId: defQTRKBillType, FDate: time.Now(), FromId: defQTRKFromId})

	if i == nil {
		return nil
	}

	i.AddModelHead(mini.QTRKHeadMini)

	i.AddModelFEntities(mini.QTRKEntityMini, mini.QTRKHeadMini.FOrgNumber)
	return i
}
