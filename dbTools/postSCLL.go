package dbTools

import (
	"fmt"
	"time"
	"xkpdaserver/jsonTools"
)

const (
	defSCLLBillType = "f4f46eb78a7149b1b7e4de98586acb67"
	defSCLLFromId   = "PRD_PickMtrl"
)

func GetPostSCLL(mini *jsonTools.SCLLMini) jsonTools.ModelBaseInterface {
	if mini == nil || mini.SCLLHeadMini == nil || mini.SCLLHeadMini.FOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := jsonTools.InitScrkModel(&jsonTools.DefModelHeadBase{FBillTypeId: defSCLLBillType, FDate: time.Now(), FromId: defSCLLFromId})

	i.AddModelHead(mini.SCLLHeadMini)

	i.AddModelFEntities(mini.SCLLEntityMini, mini.SCLLHeadMini.FOrgNumber)
	return i
}
