package dbTools

import (
	"encoding/json"
	"fmt"
)

type workShopGXJHModelBase struct {
	Formid string               `json:"formid"`
	Data   *workShopGXJHMModels `json:"data"`
}

type workShopGXJHMModels struct {
	Ids                 string            `json:"Ids"`
	Numbers             []string          `json:"Numbers"`
	EntryIds            string            `json:"EntryIds"`
	RuleId              string            `json:"RuleId"`
	TargetBillTypeId    string            `json:"TargetBillTypeId"`
	TargetFormId        string            `json:"TargetFormId"`
	IsEnableDefaultRule string            `json:"IsEnableDefaultRule"`
	IsDraftWhenSaveFail string            `json:"IsDraftWhenSaveFail"`
	CustomParams        map[string]string `json:"CustomParams"`
}

type WorkShopGXJHMini struct {
	HeadMini  *WorkShopGXJHHeadMini    `json:"HeadMini"`
	EntryMini []*WorkShopGXJHEntryMini `json:"EntityMini"`
}

type WorkShopGXJHHeadMini struct {
	FromId  string `json:"-"`
	FBillNo string `json:"FBillNo"`
	RuleId  string `json:"-"`
}

type WorkShopGXJHEntryMini struct {
	FOPERNUMBER string  `json:"FOPERNUMBER"`
	FQTY        float64 `json:"SQTY,,string" xorm:"-"`
}

var _ ModelBaseInterface = &workShopGXJHModelBase{}

func InitWorkShopModel(initBase *WorkShopGXJHMini) *workShopGXJHModelBase {
	if initBase == nil {
		return nil
	}
	return &workShopGXJHModelBase{Formid: initBase.HeadMini.FromId, Data: &workShopGXJHMModels{
		Numbers:             []string{initBase.HeadMini.FBillNo},
		RuleId:              initBase.HeadMini.RuleId,
		IsDraftWhenSaveFail: "false",
	}}
}

func (Q *workShopGXJHModelBase) CheckVerify() bool {
	return true
}

func (Q *workShopGXJHModelBase) GetJson() []byte {
	if !Q.CheckVerify() {
		fmt.Println("验证未通过")
		return nil
	}
	j, e := json.Marshal(Q)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	return j
}

func (Q *workShopGXJHModelBase) AddModelHead(in interface{}) {

}

func (Q *workShopGXJHModelBase) AddModelFEntities(ts interface{}) {

}
