package dbTools

import (
	"fmt"
	"time"
)

const (
	defWWRKBillType = "0a2c1694596d440882adb080a7a8ca1b"
	defWWRKFromId   = "STK_InStock"
)

func GetPostWWRK(mini *WWRKMini) ModelBaseInterface {
	if mini == nil || mini.HeadMini == nil || mini.HeadMini.FUseOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := InitwwrkModel(&DefModelHeadBase{FBillTypeId: defWWRKBillType, FDate: time.Now(), FromId: defWWRKFromId})

	i.AddModelHead(mini.HeadMini)

	i.AddModelFEntities(mini.EntityMini)
	return i
}
