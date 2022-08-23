package dbTools

import (
	"fmt"
	"time"
)

const (
	defSCLLBillType = "f4f46eb78a7149b1b7e4de98586acb67"
	defSCLLFromId   = "PRD_PickMtrl"
)

func GetPostSCLL(mini *SCLLMini) ModelBaseInterface {
	if mini == nil || mini.HeadMini == nil || mini.HeadMini.FUseOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := InitScllModel(&DefModelHeadBase{FBillTypeId: defSCLLBillType, FDate: time.Now(), FromId: defSCLLFromId})

	i.AddModelHead(mini.HeadMini)

	i.AddModelFEntities(mini.EntityMini)
	return i
}
