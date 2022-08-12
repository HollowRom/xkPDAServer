package dbTools

import (
	"fmt"
	"time"
)

const (
	defSCDDBillType     = "de29f16214744c21b374044d629595f2"
	defSCDDFromId       = "PRD_INSTOCK"
	defSCDDFSrcBillType = "PUR_ReceiveBill"
)

func GetPostSCRK(mini *SCRKMini) ModelBaseInterface {
	if mini == nil || mini.SCRKHeadMini == nil || mini.SCRKHeadMini.FUseOrgNumber == "" {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}
	i := InitScrkModel(&DefModelHeadBase{FBillTypeId: defSCDDBillType, FDate: time.Now(), FromId: defSCDDFromId})

	for idx := 0; idx < len(mini.SCRKEntityMini); idx++ {
		if mini.SCRKEntityMini[idx].FSrcBillType == "" {
			mini.SCRKEntityMini[idx].FSrcBillType = defSCDDFSrcBillType
		}
	}

	i.AddModelHead(mini.SCRKHeadMini)

	i.AddModelFEntities(mini.SCRKEntityMini, mini.SCRKHeadMini.FUseOrgNumber)
	return i
}
