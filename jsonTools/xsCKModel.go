package jsonTools

import (
	"encoding/json"
	"fmt"
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
		FNumber string `json:"FNumber"`
	} `json:"FSaleOrgId"`
	FCustomerID struct {
		FNumber string `json:"FNumber"`
	} `json:"FCustomerID"`
	FStockOrgId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockOrgId"`
	FSettleID struct {
		FNumber string `json:"FNumber"`
	} `json:"FSettleID"`
	FEntity []*xsckModelsEntity `json:"FInStockEntry"`
}

type xsckModelsEntity struct {
	FCustMatID struct {
		FNumber string `json:"FNumber"`
	} `json:"FCustMatID"`
	FMaterialID struct {
		FNumber string `json:"FNumber"`
	} `json:"FMaterialID"`
	FUnitID struct {
		FNumber string `json:"FNumber"`
	} `json:"FUnitID"`
	FKeeperTypeId string `json:"FKeeperTypeId"`
	FKeeperId     struct {
		FNumber string `json:"FNumber"`
	} `json:"FKeeperId"`
	FRealQty string `json:"FRealQty"`
	FPrice   string `json:"FPrice"`
	FStockID struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockID"`
	FAuxUnitQty    string `json:"FAuxUnitQty"`
	FStockStatusID struct {
		Id string `json:"Id"`
	} `json:"FStockStatusID"`
	FSalUnitID struct {
		FNumber string `json:"FNumber"`
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
	FEntity_Link_FBasePrdRealQty string `json:"FEntity_Link_FBasePrdRealQty"`
}

type XSCKMini struct {
	XSCKEntityMini []*xsckEntityMini
	XSCKHeadMini   *xsckHeadMini
}

type xsckEntityMini struct {
	FCustMatID     string
	FNumber        string
	UnitNumber     string
	FQTY           string
	FStockNumber   string
	FPrice         string
	FStockStatusId string
	FSoorDerno     string
	FSOEntryId     int
	FSOInterId     int
	FLinkInfo      []map[string]string
}

type xsckHeadMini struct {
	FOrgNumber  string
	FCustomerID string
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
	inT, ok := in.(*xsckHeadMini)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FSaleOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FCustomerID.FNumber = inT.FCustomerID
	Q.Data.Model.FStockOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FSettleID.FNumber = inT.FOrgNumber
}

func (Q *xsckModelBase) addModelFEntity(inT *xsckEntityMini, orgNumber string) {
	t := &xsckModelsEntity{
		FUnitID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FCustMatID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FCustMatID}),
		FMaterialID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FNumber}),
		FKeeperTypeId: "BD_KeeperOrg",
		FKeeperId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
		FRealQty: inT.FQTY,
		FPrice:   inT.FPrice,
		FStockID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FAuxUnitQty: inT.FQTY,
		FStockStatusID: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FSalUnitID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FSALUNITQTY: inT.FQTY,
		FSALBASEQTY: inT.FQTY,
		FSOEntryId:  inT.FSOEntryId,
		FSoorDerno:  inT.FSoorDerno,
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

func (Q *xsckModelBase) AddModelFEntities(ts interface{}, orgNumber string) {
	in, ok := ts.([]*xsckEntityMini)
	if !ok {
		return
	}
	for _, inT := range in {
		Q.addModelFEntity(inT, orgNumber)
	}
}
