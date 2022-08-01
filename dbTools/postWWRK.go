package dbTools

import (
	"fmt"
	"time"
	"xkpdaserver/jsonTools"
)

const (
	defWWRKBillType = "0a2c1694596d440882adb080a7a8ca1b"
	defWWRKFromId   = "STK_InStock"
)

func GetPostWWRK(mini *jsonTools.WWRKMini) jsonTools.ModelBaseInterface {
	if mini == nil || mini.WWRKHeadMini == nil || mini.WWRKHeadMini.FOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := jsonTools.InitwwrkModel(&jsonTools.DefModelHeadBase{FBillTypeId: defWWRKBillType, FDate: time.Now(), FromId: defWWRKFromId})

	i.AddModelHead(mini.WWRKHeadMini)

	i.AddModelFEntities(mini.WWRKEntityMini, mini.WWRKHeadMini.FOrgNumber)
	return i
}
