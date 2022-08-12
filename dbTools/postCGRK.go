package dbTools

import (
	"fmt"
	"time"
)

const (
	defCGDDBillType = "a1ff32276cd9469dad3bf2494366fa4f"
	defCGDDFromId   = "STK_InStock"
)

func GetPostCGRK(mini *CGRKMini) ModelBaseInterface {
	if mini == nil || mini.CGRKHeadMini == nil || mini.CGRKHeadMini.FUseOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := InitCgrkModel(&DefModelHeadBase{FBillTypeId: defCGDDBillType, FDate: time.Now(), FromId: defCGDDFromId})

	i.AddModelHead(mini.CGRKHeadMini)

	i.AddModelFEntities(mini.CGRKEntityMini, mini.CGRKHeadMini.FUseOrgNumber)
	return i
}
