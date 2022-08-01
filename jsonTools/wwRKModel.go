package jsonTools

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
		FNumber string `json:"FNumber"`
	} `json:"FStockOrgId"`
	FStockDeptId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockDeptId"`
	FPurchaseOrgId struct {
		FNumber string `json:"FNumber"`
	} `json:"FPurchaseOrgId"`
	FSupplierId struct {
		FNumber string `json:"FNumber"`
	} `json:"FSupplierId"`
	FSupplyId struct {
		FNumber string `json:"FNumber"`
	} `json:"FSupplyId"`
	FOwnerTypeIdHead string `json:"FOwnerTypeIdHead"`
	FOwnerIdHead     struct {
		FNumber string `json:"FNumber"`
	} `json:"FOwnerIdHead"`
	FInStockFin   *wwrkFInStockFin    `json:"FInStockFin"`
	FInStockEntry []*wwrkModelsEntity `json:"FInStockEntry"`
}

type wwrkFInStockFin struct {
	FSettleOrgId struct {
		FNumber string `json:"FNumber"`
	} `json:"FSettleOrgId"`
	FSettleTypeId struct {
		FNumber string `json:"FNumber"`
	} `json:"FSettleTypeId"`
	FSettleCurrId struct {
		FNumber string `json:"FNumber"`
	} `json:"FSettleCurrId"`
	FIsIncludedTax  string `json:"FIsIncludedTax"`
	FPriceTimePoint string `json:"FPriceTimePoint"`
}

type wwrkModelsEntity struct {
	FWWInType   string `json:"FWWInType"`
	FMaterialId struct {
		FNumber string `json:"FNumber"`
	} `json:"FMaterialId"`
	FUnitID struct {
		FNumber string `json:"FNumber"`
	} `json:"FUnitID"`
	FRealQty     string `json:"FRealQty"`
	FPriceUnitID struct {
		FNumber string `json:"FNumber"`
	} `json:"FPriceUnitID"`
	FPrice string `json:"FPrice"`
	FLot   struct {
		FNumber string `json:"FNumber"`
	} `json:"FLot"`
	FStockId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockId"`
	FDisPriceQty          string `json:"FDisPriceQty"`
	FGiveAway             string `json:"FGiveAway"`
	FOWNERTYPEID          string `json:"FOWNERTYPEID"`
	FIsReceiveUpdateStock string `json:"FIsReceiveUpdateStock"`
	FRemainInStockUnitId  struct {
		FNumber string `json:"FNumber"`
	} `json:"FRemainInStockUnitId"`
	FOWNERID struct {
		FNumber string `json:"FNumber"`
	} `json:"FOWNERID"`
	FPOOrderNo         string              `json:"FPOOrderNo"`
	FSRCBILLTYPEID     string              `json:"FSRCBILLTYPEID"`
	FSRCBillNo         string              `json:"FSRCBillNo"`
	FPOORDERENTRYID    int                 `json:"FPOORDERENTRYID"`
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
	WWRKEntityMini []*wwrkEntityMini
	WWRKHeadMini   *wwrkHeadMini
}

type wwrkEntityMini struct {
	FParentNumber   string
	FNumber         string
	UnitNumber      string
	FQTY            string
	FPrice          string
	FStockNumber    string
	FStockStatusId  string
	FSrcInterId     int
	FSrcEntryId     int
	FSrcBillNo      string
	FSrcEntrySeq    int
	FSubReqId       int
	FPPbomBillNo    string
	FSubReqBillNo   string
	FPPbomEntryId   int
	FPOOrderBillNo  string
	FPOORDERINTERID int
	FPOORDERENTRYID int
	FKeeperId       string
	FLotNo          string
	FLinkInfo       []map[string]string
}

type wwrkHeadMini struct {
	FBillNo         string
	FOrgNumber      string
	FSUPPLIERNumber string
	FSUPPLIERName   string
}

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
	inT, ok := in.(*wwrkHeadMini)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FStockOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FStockDeptId.FNumber = inT.FOrgNumber
	Q.Data.Model.FPurchaseOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FSupplierId.FNumber = inT.FSUPPLIERNumber
	Q.Data.Model.FSupplyId.FNumber = inT.FSUPPLIERNumber
	Q.Data.Model.FOwnerTypeIdHead = "BD_OwnerOrg"
	Q.Data.Model.FOwnerIdHead.FNumber = inT.FSUPPLIERNumber
	Q.Data.Model.FInStockFin = &wwrkFInStockFin{
		FSettleOrgId: struct {
			FNumber string `json:"FNumber"`
		}{inT.FOrgNumber},
		FSettleTypeId: struct {
			FNumber string `json:"FNumber"`
		}{"PRE001"},
		FSettleCurrId: struct {
			FNumber string `json:"FNumber"`
		}{"PRE001"},
		FIsIncludedTax:  "false",
		FPriceTimePoint: "1",
	}
}

func (Q *wwrkModelBase) addModelFEntity(inT *wwrkEntityMini, orgNumber string) {
	t := &wwrkModelsEntity{
		FWWInType: "QLI",
		FMaterialId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FNumber}),
		FUnitID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FRealQty: inT.FQTY,
		FPriceUnitID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FPrice: inT.FPrice,
		FLot: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FLotNo}),
		FStockId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FOWNERTYPEID: "BD_OwnerOrg",
		FRemainInStockUnitId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FOWNERID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
		FPOOrderNo:      inT.FPOOrderBillNo,
		FSRCBILLTYPEID:  "PUR_PurchaseOrder",
		FSRCBillNo:      inT.FSrcBillNo,
		FPOORDERENTRYID: inT.FPOORDERENTRYID,
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

func (Q *wwrkModelBase) AddModelFEntities(ts interface{}, orgNumber string) {
	in, ok := ts.([]*wwrkEntityMini)
	if !ok {
		return
	}
	for _, inT := range in {
		Q.addModelFEntity(inT, orgNumber)
	}
}
