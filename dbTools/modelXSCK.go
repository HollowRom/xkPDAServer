package dbTools

import (
	"encoding/json"
	"fmt"
	"strings"
)

type xsckModelBase struct {
	Formid string         `json:"formid"`
	Data   *xsckModelData `json:"data"`
}

type xsckModelData struct {
	Model *xsckModels `json:"Model"`
}

type xsckModels struct {
	FBillTypeID struct {
		Id string `json:"Id"`
	} `json:"FBillTypeID"`
	FDate      string `json:"FDate"`
	FSaleOrgId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSaleOrgId"`
	FCustomerID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FCustNumber"`
	FStockOrgId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockOrgId"`
	FSettleID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSettleID"`
	FEntity []*xsckModelsEntity `json:"FInStockEntry"`
}

type xsckModelsEntity struct {
	FCustMatID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FCustNumber"`
	FMaterialID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FMaterialID"`
	FUnitID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FUnitID"`
	FKeeperTypeId string `json:"FKeeperTypeId"`
	FKeeperId     struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FKeeperId"`
	FRealQty string `json:"FRealQty"`
	FPrice   string `json:"FPrice"`
	FLot struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FLOT_TEXT"`
	FStockID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockID"`
	FAuxUnitQty    string `json:"FAuxUnitQty"`
	FStockStatusID struct {
		Id string `json:"Id"`
	} `json:"FStockStatusID"`
	FSalUnitID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSalUnitID"`
	FSALUNITQTY  string              `json:"FSALUNITQTY"`
	FSALBASEQTY  string              `json:"FSALBASEQTY"`
	FSOEntryId   int                 `json:"FENTRYID"`
	FSoorDerno   string              `json:"FBILLNO"`
	FEntity_Link []*xsckFEntity_Link `json:"FEntity_Link"`
}

type xsckFEntity_Link struct {
	FEntity_Link_FRuleId         string `json:"FEntity_Link_FRuleId"`
	FEntity_Link_FSTableName     string `json:"FEntity_Link_FSTableName"`
	FEntity_Link_FSBillId        string `json:"FEntity_Link_FSBillId"`
	FEntity_Link_FSId            string `json:"FEntity_Link_FSId"`
	FEntity_Link_FBasePrdRealQty string `json:"FEntity_Link_FBasePrdRealQty"`
}

type XSCKMini struct {
	EntityMini []*XSDDEntry
	HeadMini   *XSDDMain
}

//type xsckEntityMini struct {
//	FCustNumber string
//	Id         string
//	FBaseUnitNumber string
//	FMustQty        string
//	FStockNumber    string
//	FPrice         string
//	FStockStatusId string
//	FBILLNO    string
//	FENTRYID  int
//	FID       int
//	FLinkInfo []map[string]string
//}

//type xsckHeadMini struct {
//	FUseOrgNumber string
//	FCustNumber   string
//}

var _ ModelBaseInterface = &xsckModelBase{}

func InitxsckModel(initBase *DefModelHeadBase) *xsckModelBase {
	if initBase == nil {
		return nil
	}
	return &xsckModelBase{Formid: initBase.FromId, Data: &xsckModelData{
		Model: &xsckModels{
			FBillTypeID: struct {
				Id string `json:"Id"`
			}{Id: initBase.FBillTypeId},
			FDate: initBase.FDate.Format("2006-01-02"),
		},
	}}
}

func (Q *xsckModelBase) CheckVerify() bool {
	return true
}

func (Q *xsckModelBase) GetJson() []byte {
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

func (Q *xsckModelBase) AddModelHead(in interface{}) {
	inT, ok := in.(*XSDDMain)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FSaleOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FCustomerID.FNumber = inT.FCustNumber
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FSettleID.FNumber = inT.FUseOrgNumber
}

func (Q *xsckModelBase) addModelFEntity(inT *XSDDEntry) {
	t := &xsckModelsEntity{
		FUnitID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FCustMatID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FCustNumber}),
		FMaterialID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FNUMBER}),
		FKeeperTypeId: "BD_KeeperOrg",
		FKeeperId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FUseOrgNumber}),
		FRealQty: inT.FMustQty,
		FPrice:   inT.FPrice,
		FLot: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: strings.TrimRight(inT.FLOT_TEXT, " ")}),
		FStockID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FAuxUnitQty: inT.FMustQty,
		FStockStatusID: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FSalUnitID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FSALUNITQTY: inT.FMustQty,
		FSALBASEQTY: inT.FMustQty,
		FSOEntryId:  inT.FENTRYID,
		FSoorDerno:  inT.FBILLNO,
	}
	t.FEntity_Link = append(t.FEntity_Link, &xsckFEntity_Link{
		FEntity_Link_FRuleId:         inT.FLinkInfo[0]["FEntity_Link_FRuleId"],
		FEntity_Link_FSTableName:     inT.FLinkInfo[0]["FEntity_Link_FSTableName"],
		FEntity_Link_FSBillId:        inT.FLinkInfo[0]["FEntity_Link_FSBillId"],
		FEntity_Link_FSId:            inT.FLinkInfo[0]["FEntity_Link_FSId"],
		FEntity_Link_FBasePrdRealQty: inT.FLinkInfo[0]["FEntity_Link_FBasePrdRealQty"],
	})
	Q.Data.Model.FEntity = append(Q.Data.Model.FEntity, t)
}

func (Q *xsckModelBase) AddModelFEntities(ts interface{}) {
	in, ok := ts.([]*XSDDEntry)
	if !ok {
		return
	}
	for _, inT := range in {
		Q.addModelFEntity(inT)
	}
}
