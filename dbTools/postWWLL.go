package dbTools

import (
	"fmt"
	"time"
)

const (
	defWWLLBillType = "4518706ee0e84af49671ba2af1498d48"
	defWWLLFromId   = "SUB_PickMtrl"
)

func GetPostWWLL(mini *WWLLMini) ModelBaseInterface {
	if mini == nil || mini.HeadMini == nil || mini.HeadMini.FUseOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := InitwwllModel(&DefModelHeadBase{FBillTypeId: defWWLLBillType, FDate: time.Now(), FromId: defWWLLFromId})

	i.AddModelHead(mini.HeadMini)

	i.AddModelFEntities(mini.EntityMini)
	return i
}
