package dbTools

import (
	"fmt"
	"time"
	"xkpdaserver/jsonTools"
)

const (
	defCGDDBillType = "a1ff32276cd9469dad3bf2494366fa4f"
	defCGDDFromId   = "STK_InStock"
)

func GetPostCGRK(mini *jsonTools.CGRKMini) jsonTools.ModelBaseInterface {
	if mini == nil || mini.CGRKHeadMini == nil || mini.CGRKHeadMini.FOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := jsonTools.InitCgrkModel(&jsonTools.DefModelHeadBase{FBillTypeId: defCGDDBillType, FDate: time.Now(), FromId: defCGDDFromId})

	i.AddModelHead(mini.CGRKHeadMini)

	i.AddModelFEntities(mini.CGRKEntityMini, mini.CGRKHeadMini.FOrgNumber)
	return i
}
