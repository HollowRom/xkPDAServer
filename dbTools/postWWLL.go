package dbTools

import (
	"fmt"
	"time"
	"xkpdaserver/jsonTools"
)

const (
	defWWLLBillType = "4518706ee0e84af49671ba2af1498d48"
	defWWLLFromId   = "SUB_PickMtrl"
)

func GetPostWWLL(mini *jsonTools.WWLLMini) jsonTools.ModelBaseInterface {
	if mini == nil || mini.WWLLHeadMini == nil || mini.WWLLHeadMini.FOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := jsonTools.InitwwllModel(&jsonTools.DefModelHeadBase{FBillTypeId: defWWLLBillType, FDate: time.Now(), FromId: defWWLLFromId})

	i.AddModelHead(mini.WWLLHeadMini)

	i.AddModelFEntities(mini.WWLLEntityMini, mini.WWLLHeadMini.FOrgNumber)
	return i
}
