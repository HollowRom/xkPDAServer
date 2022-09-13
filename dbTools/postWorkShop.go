package dbTools

import (
	"fmt"
)

const (
	defWorkShopRuleId = "SFC_OPRTPLAN2OPTRRPT"
	defWorkShopFromId = "SFC_OperationPlanning"
)

func GetPostWorkShop(mini *WorkShopGXJHMini) ModelBaseInterface {
	if mini == nil || mini.HeadMini == nil || mini.HeadMini.FBillNo == "" || mini.EntryMini == nil || len(mini.EntryMini) == 0 {
		fmt.Println("输入mini缺少必须的数据")
		return nil
	}

	if mini.HeadMini.FromId == "" {
		mini.HeadMini.FromId = defWorkShopFromId
	}

	if mini.HeadMini.RuleId == "" {
		mini.HeadMini.RuleId = defWorkShopRuleId
	}

	i := InitWorkShopModel(mini)

	return i
}
