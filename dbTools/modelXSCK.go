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
	} `json:"FCustomerID"`
	FStockerID struct{
		FNumber string `json:"FNumber"`
	}
	FStockOrgId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockOrgId"`
	FSettleID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSettleID"`
	FEntity []*xsckModelsEntity `json:"FEntity"`
}

type xsckModelsEntity struct {
	FCustMatID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FCustMatID"`
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
	FSOEntryId   int                 `json:"FSOEntryId"`
	FSoorDerno   string              `json:"FSoorDerno"`
	FEntity_Link []*xsckFEntity_Link `json:"FEntity_Link"`
}

type xsckFEntity_Link struct {
	FEntity_Link_FRuleId         string `json:"FEntity_Link_FRuleId"`
	FEntity_Link_FSTableName     string `json:"FEntity_Link_FSTableName"`
	FEntity_Link_FSBillId        string `json:"FEntity_Link_FSBillId"`
	FEntity_Link_FSId            string `json:"FEntity_Link_FSId"`
	FEntity_Link_FBaseUnitQty string `json:"FEntity_Link_FBaseUnitQty"`
	FEntity_Link_FSALBASEQTY string `json:"FEntity_Link_FSALBASEQTY"`
}

type XSCKMini struct {
	EntityMini []*XSDDEntry
}

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
	inT, ok := in.(*XSDDEntry)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FSaleOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FCustomerID.FNumber = inT.FCustNumber
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FSettleID.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FStockerID.FNumber = inT.FStockerId //"027"
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
		FRealQty: inT.SQTY,
		FPrice:   inT.FPrice,
		FLot: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: strings.TrimRight(inT.FLOT_TEXT, " ")}),
		FStockID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FAuxUnitQty: inT.SQTY,
		FStockStatusID: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FSalUnitID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FSALUNITQTY: inT.SQTY,
		FSALBASEQTY: inT.SQTY,
		FSOEntryId:  inT.FSOEntryId,
		FSoorDerno:  inT.FORDERNO,
	}
	t.FEntity_Link = append(t.FEntity_Link, &xsckFEntity_Link{
		FEntity_Link_FRuleId:         inT.FLinkInfo[0]["FEntity_Link_FRuleId"],
		FEntity_Link_FSTableName:     inT.FLinkInfo[0]["FEntity_Link_FSTableName"],
		FEntity_Link_FSBillId:        inT.FLinkInfo[0]["FEntity_Link_FSBillId"],
		FEntity_Link_FSId:            inT.FLinkInfo[0]["FEntity_Link_FSId"],
		FEntity_Link_FBaseUnitQty: inT.FLinkInfo[0]["FEntity_Link_FBaseUnitQty"],
		FEntity_Link_FSALBASEQTY: inT.FLinkInfo[0]["FEntity_Link_FSALBASEQTY"],
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
