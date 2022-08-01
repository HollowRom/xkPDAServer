package jsonTools

import (
	"encoding/json"
	"fmt"
)

type cgrkModelBase struct {
	Formid string         `json:"formid"`
	Data   *cgrkModelData `json:"data"`
}

type cgrkModelData struct {
	Model *cgrkModels `json:"Model"`
}

type cgrkModels struct {
	FBillTypeID struct {
		Id string `json:"Id"`
	} `json:"FBillTypeID"`
	FDate       string `json:"FDate"`
	FStockOrgId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockOrgId"`
	FSupplierId struct {
		FNumber string `json:"FNumber"`
	} `json:"FSupplierId"`
	FInStockEntry []*cgrkModelsEntity `json:"FInStockEntry"`
}

type cgrkModelsEntity struct {
	FMaterialId struct {
		FNumber string `json:"FNumber"`
	} `json:"FMaterialId"`
	FUnitID struct {
		FNumber string `json:"FNumber"`
	} `json:"FUnitID"`
	FRealQty string `json:"FRealQty"`
	FPrice   string `json:"FPrice"`
	FStockId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockId"`
	FOWNERID struct {
		FNumber string `json:"FNumber"`
	} `json:"FOWNERID"`
	FPOOrderNo        string                    `json:"FPOOrderNo"`
	FSRCBILLTYPEID    string                    `json:"FSRCBILLTYPEID"`
	FSRCBillNo        string                    `json:"FSRCBillNo"`
	FPOORDERENTRYID   int                       `json:"FPOORDERENTRYID"`
	FinstockentryLink []*cgrkFInStockEntry_Link `json:"FInStockEntry_Link"`
}

type cgrkFInStockEntry_Link struct {
	FInStockEntryLinkFRuleId               string `json:"FInStockEntry_Link_FRuleId"`
	FInStockEntryLinkFSTableName           string `json:"FInStockEntry_Link_FSTableName"`
	FInStockEntryLinkFSBillId              string `json:"FInStockEntry_Link_FSBillId"`
	FInStockEntryLinkFSId                  string `json:"FInStockEntry_Link_FSId"`
	FInStockEntryLinkFBaseUnitQty          string `json:"FInStockEntry_Link_FBaseUnitQty"`
	FInStockEntryLinkFRemainInStockBaseQty string `json:"FInStockEntry_Link_FRemainInStockBaseQty"`
}

type CGRKMini struct {
	CGRKEntityMini []*CGRKEntityMini
	CGRKHeadMini   *CGRKHeadMini
}

type CGRKEntityMini struct {
	FNumber        string
	UnitNumber     string
	FQTY           string
	FPrice         string
	FStockNumber   string
	FNote          string
	FStockStatusId string
	FKeeperId      string
	FLotNo         string
	FOrderNo       string
	FSrcBillNo     string
	FSrcBillType   string
	FOrderInterId  int
	FOrderEntryId  int
	FLinkInfo      []map[string]string
}

type CGRKHeadMini struct {
	FOrgNumber string
	FSupplyId  string
}

var _ ModelBaseInterface = &cgrkModelBase{}

func InitCgrkModel(initBase *DefModelHeadBase) *cgrkModelBase {
	if initBase == nil {
		return nil
	}
	return &cgrkModelBase{Formid: initBase.FromId, Data: &cgrkModelData{
		Model: &cgrkModels{
			FBillTypeID: struct {
				Id string `json:"Id"`
			}(struct{ Id string }{Id: initBase.FBillTypeId}),
			FDate: initBase.FDate.Format("2006-01-02"),
		},
	}}
}

func (Q *cgrkModelBase) CheckVerify() bool {
	return true
}

func (Q *cgrkModelBase) GetJson() []byte {
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

func (Q *cgrkModelBase) AddModelHead(in interface{}) {
	inT, ok := in.(*CGRKHeadMini)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FSupplierId.FNumber = inT.FSupplyId
}

func (Q *cgrkModelBase) addModelFEntity(inT *CGRKEntityMini, orgNumber string) {
	t := &cgrkModelsEntity{
		FMaterialId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FNumber}),
		FUnitID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FRealQty: inT.FQTY,
		FPrice:   inT.FPrice,
		FStockId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FOWNERID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
		FPOOrderNo:      inT.FOrderNo,
		FSRCBILLTYPEID:  inT.FSrcBillType,
		FSRCBillNo:      inT.FSrcBillNo,
		FPOORDERENTRYID: inT.FOrderEntryId,
	}
	if inT.FLinkInfo != nil && len(inT.FLinkInfo) == 1 {
		tempLinkMap := &cgrkFInStockEntry_Link{
			FInStockEntryLinkFRuleId:               inT.FLinkInfo[0]["FInStockEntry_Link_FRuleId"],
			FInStockEntryLinkFSTableName:           inT.FLinkInfo[0]["FInStockEntry_Link_FSTableName"],
			FInStockEntryLinkFSBillId:              inT.FLinkInfo[0]["FInStockEntry_Link_FSBillId"],
			FInStockEntryLinkFSId:                  inT.FLinkInfo[0]["FInStockEntry_Link_FSId"],
			FInStockEntryLinkFBaseUnitQty:          inT.FLinkInfo[0]["FInStockEntry_Link_FBaseUnitQty"],
			FInStockEntryLinkFRemainInStockBaseQty: inT.FLinkInfo[0]["FInStockEntry_Link_FRemainInStockBaseQty"],
		}
		t.FinstockentryLink = append(t.FinstockentryLink, tempLinkMap)
	}
	Q.Data.Model.FInStockEntry = append(Q.Data.Model.FInStockEntry, t)
}

func (Q *cgrkModelBase) AddModelFEntities(ts interface{}, orgNumber string) {
	ins, ok := ts.([]*CGRKEntityMini)
	if !ok {
		return
	}
	for _, inT := range ins {
		Q.addModelFEntity(inT, orgNumber)
	}
}
