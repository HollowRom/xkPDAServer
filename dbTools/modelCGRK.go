package dbTools

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
		FNumber string `json:"FNUMBER"`
	} `json:"FStockOrgId"`
	FSupplierId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSupplierId"`
	FInStockEntry []*cgrkModelsEntity `json:"FInStockEntry"`
}

type cgrkModelsEntity struct {
	FMaterialId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FMaterialId"`
	FUnitID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FUnitID"`
	FRealQty string `json:"FRealQty"`
	FPrice   string `json:"FPrice"`
	FStockId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockId"`
	FOWNERID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FOWNERID"`
	FPOOrderNo        string                    `json:"FPOOrderNo"`
	FSRCBILLTYPEID    string                    `json:"FSRCBILLTYPEID"`
	FSRCBillNo        string                    `json:"FSRCBillNo"`
	FPOORDERENTRYID   int                       `json:"FENTRYID"`
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
	EntityMini []*CGDDEntry
	HeadMini   *CGDDMain
}

//type EntityMini struct {
//	FNUMBER         string
//	FBaseUnitNumber string
//	FMustQty        string
//	FPrice          string //**
//	FStockNumber    string
//	FNote           string
//	FStockStatusId  string
//	FKeeperId       string
//	FLOT_TEXT  string
//	FSRCBILLNO string
//	FBILLNO    string
//	FSrcBillType  string
//	FSRCID      int
//	FSRCENTRYID int
//	FLinkInfo   []map[string]string
//}

//type HeadMini struct {
//	FUseOrgNumber string
//	FSupplierNumber   string
//}

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
	inT, ok := in.(*CGDDMain)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FSupplierId.FNumber = inT.FSuppNumber
}

func (Q *cgrkModelBase) addModelFEntity(inT *CGDDEntry) {
	t := &cgrkModelsEntity{
		FMaterialId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FNUMBER}),
		FUnitID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FRealQty: inT.FMustQty,
		FPrice:   inT.FPrice,
		FStockId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FOWNERID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FUseOrgNumber}),
		FPOOrderNo:      inT.FSRCBILLNO,
		FSRCBILLTYPEID:  inT.FSrcBillType,
		FSRCBillNo:      inT.FBILLNO,
		FPOORDERENTRYID: inT.FSRCENTRYID,
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

func (Q *cgrkModelBase) AddModelFEntities(ts interface{}) {
	ins, ok := ts.([]*CGDDEntry)
	if !ok {
		return
	}
	for _, inT := range ins {
		Q.addModelFEntity(inT)
	}
}
