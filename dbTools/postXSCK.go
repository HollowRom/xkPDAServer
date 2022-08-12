package dbTools

import (
	"fmt"
	"time"
)

const (
	defXSCKBillType = "ad0779a4685a43a08f08d2e42d7bf3e9"
	defXSCKFromId   = "SAL_OUTSTOCK"
)

func GetPostXSCK(mini *XSCKMini) ModelBaseInterface {
	if mini == nil || mini.XSCKHeadMini == nil || mini.XSCKHeadMini.FOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := InitScrkModel(&DefModelHeadBase{FBillTypeId: defXSCKBillType, FDate: time.Now(), FromId: defXSCKFromId})

	i.AddModelHead(mini.XSCKHeadMini)

	i.AddModelFEntities(mini.XSCKEntityMini, mini.XSCKHeadMini.FOrgNumber)
	return i
}
