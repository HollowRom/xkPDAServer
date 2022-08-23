package dbTools

import (
	"encoding/json"
	"fmt"
)

type wwrkModelBase struct {
	Formid string         `json:"formid"`
	Data   *wwrkModelData `json:"data"`
}

type wwrkModelData struct {
	Model *wwrkModels `json:"Model"`
}

type wwrkModels struct {
	FBillTypeID struct {
		Id string `json:"Id"`
	} `json:"FBillTypeID"`
	FDate       string `json:"FDate"`
	FStockOrgId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockOrgId"`
	FStockDeptId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockDeptId"`
	FPurchaseOrgId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FPurchaseOrgId"`
	FSupplierId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSupplierId"`
	FSupplyId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSupplierNumber"`
	FOwnerTypeIdHead string `json:"FOwnerTypeIdHead"`
	FOwnerIdHead     struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FOwnerIdHead"`
	FInStockFin   *wwrkFInStockFin    `json:"FInStockFin"`
	FInStockEntry []*wwrkModelsEntity `json:"FInStockEntry"`
}

type wwrkFInStockFin struct {
	FSettleOrgId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSettleOrgId"`
	FSettleTypeId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSettleTypeId"`
	FSettleCurrId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSettleCurrId"`
	FIsIncludedTax  string `json:"FIsIncludedTax"`
	FPriceTimePoint string `json:"FPriceTimePoint"`
}

type wwrkModelsEntity struct {
	FWWInType   string `json:"FWWInType"`
	FMaterialId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FMaterialId"`
	FUnitID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FUnitID"`
	FRealQty     string `json:"FRealQty"`
	FPriceUnitID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FPriceUnitID"`
	FPrice string `json:"FPrice"`
	FLot   struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FLOT_TEXT"`
	FStockId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockId"`
	FDisPriceQty          string `json:"FDisPriceQty"`
	FGiveAway             string `json:"FGiveAway"`
	FOWNERTYPEID          string `json:"FOWNERTYPEID"`
	FIsReceiveUpdateStock string `json:"FIsReceiveUpdateStock"`
	FRemainInStockUnitId  struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FRemainInStockUnitId"`
	FOWNERID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FOWNERID"`
	FPOOrderNo         string              `json:"FPOOrderNo"`
	FSRCBILLTYPEID     string              `json:"FSRCBILLTYPEID"`
	FSRCBillNo         string              `json:"FSRCBillNo"`
	FPOORDERENTRYID    int                 `json:"FENTRYID"`
	FInStockEntry_Link []*wwrkFEntity_Link `json:"FInStockEntry_Link"`
}

type wwrkFEntity_Link struct {
	FInStockEntryLinkFRuleId        string `json:"FInStockEntry_Link_FRuleId"`
	FInStockEntryLinkFSTableName    string `json:"FInStockEntry_Link_FSTableName"`
	FInStockEntryLinkFSBillId       string `json:"FInStockEntry_Link_FSBillId"`
	FInStockEntryLinkFSId           string `json:"FInStockEntry_Link_FSId"`
	FInStockEntryLinkFBaseActualQty string `json:"FInStockEntry_Link_FBaseActualQty"`
}

type WWRKMini struct {
	EntityMini []*WWDDEntry
	HeadMini   *WWDDMain
}

//type wwrkEntityMini struct {
//	FNUMBER         string
//	FBaseUnitNumber string
//	FMustQty        string
//	FPrice          string
//	FStockNumber    string
//	FStockStatusId  string
//	FSrcBillNo      string //收料通知单单号
//	FBILLNO         string
//	FID       int
//	FENTRYID  int
//	FKeeperId string
//	FLOT_TEXT string
//	FLinkInfo []map[string]string
//}

//type wwrkHeadMini struct {
//	FBILLNO         string
//	FUseOrgNumber   string
//	FSupplierNumber string
//	FSupplierName   string
//}

var _ ModelBaseInterface = &wwrkModelBase{}

func InitwwrkModel(initBase *DefModelHeadBase) *wwrkModelBase {
	if initBase == nil {
		return nil
	}
	return &wwrkModelBase{Formid: initBase.FromId, Data: &wwrkModelData{
		Model: &wwrkModels{
			FBillTypeID: struct {
				Id string `json:"Id"`
			}{Id: initBase.FBillTypeId},
			FDate: initBase.FDate.Format("2006-01-02"),
		},
	}}
}

func (Q *wwrkModelBase) CheckVerify() bool {
	return true
}

func (Q *wwrkModelBase) GetJson() []byte {
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

func (Q *wwrkModelBase) AddModelHead(in interface{}) {
	inT, ok := in.(*WWDDMain)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FStockDeptId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FPurchaseOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FSupplierId.FNumber = inT.FSupplierNumber
	Q.Data.Model.FSupplyId.FNumber = inT.FSupplierNumber
	Q.Data.Model.FOwnerTypeIdHead = "BD_OwnerOrg"
	Q.Data.Model.FOwnerIdHead.FNumber = inT.FSupplierNumber
	Q.Data.Model.FInStockFin = &wwrkFInStockFin{
		FSettleOrgId: struct {
			FNumber string `json:"FNUMBER"`
		}{inT.FUseOrgNumber},
		FSettleTypeId: struct {
			FNumber string `json:"FNUMBER"`
		}{"PRE001"},
		FSettleCurrId: struct {
			FNumber string `json:"FNUMBER"`
		}{"PRE001"},
		FIsIncludedTax:  "false",
		FPriceTimePoint: "1",
	}
}

func (Q *wwrkModelBase) addModelFEntity(inT *WWDDEntry) {
	t := &wwrkModelsEntity{
		FWWInType: "QLI",
		FMaterialId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FNUMBER}),
		FUnitID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FRealQty: inT.FMustQty,
		FPriceUnitID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FPrice: inT.FPrice,
		FLot: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FLOT_TEXT}),
		FStockId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FOWNERTYPEID: "BD_OwnerOrg",
		FRemainInStockUnitId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FOWNERID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FUseOrgNumber}),
		FPOOrderNo:      inT.FBILLNO,
		FSRCBILLTYPEID:  "PUR_PurchaseOrder",
		FSRCBillNo:      inT.FSrcBillNo,
		FPOORDERENTRYID: inT.FENTRYID,
	}
	t.FInStockEntry_Link = append(t.FInStockEntry_Link, &wwrkFEntity_Link{
		FInStockEntryLinkFRuleId:        inT.FLinkInfo[0]["FInStockEntry_Link_FRuleId"],
		FInStockEntryLinkFSTableName:    inT.FLinkInfo[0]["FInStockEntry_Link_FSTableName"],
		FInStockEntryLinkFSBillId:       inT.FLinkInfo[0]["FInStockEntry_Link_FSBillId"],
		FInStockEntryLinkFSId:           inT.FLinkInfo[0]["FInStockEntry_Link_FSId"],
		FInStockEntryLinkFBaseActualQty: inT.FLinkInfo[0]["FInStockEntry_Link_FBaseActualQty"],
	})
	Q.Data.Model.FInStockEntry = append(Q.Data.Model.FInStockEntry, t)
}

func (Q *wwrkModelBase) AddModelFEntities(ts interface{}) {
	in, ok := ts.([]*WWDDEntry)
	if !ok {
		return
	}
	for _, inT := range in {
		Q.addModelFEntity(inT)
	}
}
