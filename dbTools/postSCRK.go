package dbTools

import (
	"fmt"
	"time"
	"xkpdaserver/jsonTools"
)

const (
	defSCDDBillType     = "de29f16214744c21b374044d629595f2"
	defSCDDFromId       = "PRD_INSTOCK"
	defSCDDFSrcBillType = "PUR_ReceiveBill"
)

func GetPostSCRK(mini *jsonTools.SCRKMini) jsonTools.ModelBaseInterface {
	if mini == nil || mini.SCRKHeadMini == nil || mini.SCRKHeadMini.FOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := jsonTools.InitScrkModel(&jsonTools.DefModelHeadBase{FBillTypeId: defSCDDBillType, FDate: time.Now(), FromId: defSCDDFromId})

	for idx := 0; idx < len(mini.SCRKEntityMini); idx++ {
		if mini.SCRKEntityMini[idx].FSrcBillType == "" {
			mini.SCRKEntityMini[idx].FSrcBillType = defSCDDFSrcBillType
		}
	}

	i.AddModelHead(mini.SCRKHeadMini)

	i.AddModelFEntities(mini.SCRKEntityMini, mini.SCRKHeadMini.FOrgNumber)
	return i
}
