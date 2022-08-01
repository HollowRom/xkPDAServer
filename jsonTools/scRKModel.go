package jsonTools

import (
	"encoding/json"
	"fmt"
)

type scrkModelBase struct {
	Formid string         `json:"formid"`
	Data   *scrkModelData `json:"data"`
}

type scrkModelData struct {
	Model *scrkModels `json:"Model"`
}

type scrkModels struct {
	FBillType struct {
		Id string `json:"Id"`
	} `json:"FBillType"`
	FDate       string `json:"FDate"`
	FStockOrgId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockOrgId"`
	FPrdOrgId struct {
		FNumber string `json:"FNumber"`
	} `json:"FPrdOrgId"`
	FOwnerTypeId string `json:"FOwnerTypeId"`
	FOwnerId0    struct {
		FNumber string `json:"FNumber"`
	} `json:"FOwnerId0"`
	FEntity []*scrkModelsEntity `json:"FInStockEntry"`
}

type scrkModelsEntity struct {
	FSrcEntryId int `json:"FSrcEntryId"`
	FMaterialId struct {
		FNumber string `json:"FNumber"`
	} `json:"FMaterialId"`
	FUnitID struct {
		FNumber string `json:"FNumber"`
	} `json:"FUnitID"`
	FMustQty    string `json:"FMustQty"`
	FRealQty    string `json:"FRealQty"`
	FBaseUnitId struct {
		FNumber string `json:"FNumber"`
	} `json:"FBaseUnitId"`
	FOwnerTypeId string `json:"FOwnerTypeId"`
	FOwnerId     struct {
		FNumber string `json:"FNumber"`
	} `json:"FOwnerId"`
	FStockId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockId"`
	FLot struct {
		FNumber string `json:"FNumber"`
	} `json:"FLot"`
	FMoBillNo    string `json:"FMoBillNo"`
	FMoId        int    `json:"FMoId"`
	FMoEntryId   int    `json:"FMoEntryId"`
	FMoEntrySeq  int    `json:"FMoEntrySeq"`
	FStockUnitId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockUnitId"`
	FSrcBillType   string `json:"FSrcBillType"`
	FSrcBillNo     string `json:"FSrcBillNo"`
	FSrcInterId    int    `json:"FSrcInterId"`
	FStockStatusId struct {
		Id string `json:"Id"`
	} `json:"FStockStatusId"`
	FSrcEntrySeq   int    `json:"FSrcEntrySeq"`
	FMOMAINENTRYID int    `json:"FMOMAINENTRYID"`
	FKeeperTypeId  string `json:"FKeeperTypeId"`
	FKeeperId      struct {
		FNumber string `json:"FNumber"`
	} `json:"FKeeperId"`
	FentityLink []*scrkFENTITY_Link `json:"FENTITY_Link"`
}

type scrkFENTITY_Link struct {
	FENTITYLinkFRuleId        string `json:"FENTITY_Link_FRuleId"`
	FENTITYLinkFSTableName    string `json:"FENTITY_Link_FSTableName"`
	FENTITYLinkFSBillId       string `json:"FENTITY_Link_FSBillId"`
	FENTITYLinkFSId           string `json:"FENTITY_Link_FSId"`
	FENTITYLinkFBaseActualQty string `json:"FENTITY_Link_FBaseActualQty"`
}

type SCRKMini struct {
	SCRKEntityMini []*scrkEntityMini
	SCRKHeadMini   *scrkHeadMini
}

type scrkEntityMini struct {
	FNumber        string
	FSrcEntryId    int
	UnitNumber     string
	FQTY           string
	FPrice         string
	FStockNumber   string
	FNote          string
	FStockStatusId string
	FLot           string
	FMoBillNo      string
	FMoId          int
	FMoEntryId     int
	FMoEntrySeq    int
	FSrcBillNo     string
	FSrcBillType   string
	FSrcInterId    int
	FLinkInfo      []map[string]string
}

type scrkHeadMini struct {
	FOrgNumber string
}

var _ ModelBaseInterface = &scrkModelBase{}

func InitScrkModel(initBase *DefModelHeadBase) *scrkModelBase {
	if initBase == nil {
		return nil
	}
	return &scrkModelBase{Formid: initBase.FromId, Data: &scrkModelData{
		Model: &scrkModels{
			FBillType: struct {
				Id string `json:"Id"`
			}(struct{ Id string }{Id: initBase.FBillTypeId}),
			FDate: initBase.FDate.Format("2006-01-02"),
		},
	}}
}

func (Q *scrkModelBase) CheckVerify() bool {
	return true
}

func (Q *scrkModelBase) GetJson() []byte {
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

func (Q *scrkModelBase) AddModelHead(in interface{}) {
	inT, ok := in.(*scrkHeadMini)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FPrdOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FPrdOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FOwnerTypeId = "BD_OwnerOrg"
	Q.Data.Model.FOwnerId0.FNumber = inT.FOrgNumber
}

func (Q *scrkModelBase) addModelFEntity(inT *scrkEntityMini, orgNumber string) {
	t := &scrkModelsEntity{
		FSrcEntryId: inT.FSrcEntryId,
		FMaterialId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FNumber}),
		FUnitID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FMustQty: inT.FQTY,
		FBaseUnitId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FRealQty:     inT.FQTY,
		FOwnerTypeId: "BD_OwnerOrg",
		FOwnerId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
		FStockId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FLot: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FLot}),
		FStockUnitId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FStockStatusId: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FKeeperTypeId: "BD_KeeperOrg",
		FKeeperId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
	}

	if inT.FLinkInfo != nil && len(inT.FLinkInfo) == 1 {
		t.FentityLink = append(t.FentityLink, &scrkFENTITY_Link{
			FENTITYLinkFRuleId:        inT.FLinkInfo[0]["FENTITY_Link_FRuleId"],
			FENTITYLinkFSTableName:    inT.FLinkInfo[0]["FENTITY_Link_FSTableName"],
			FENTITYLinkFSBillId:       inT.FLinkInfo[0]["FENTITY_Link_FSBillId"],
			FENTITYLinkFSId:           inT.FLinkInfo[0]["FENTITY_Link_FSId"],
			FENTITYLinkFBaseActualQty: inT.FLinkInfo[0]["FENTITY_Link_FBaseActualQty"],
		})
		t.FMoBillNo = inT.FMoBillNo
		t.FMoId = inT.FMoId
		t.FMoEntrySeq = inT.FMoEntrySeq
		t.FSrcEntryId = inT.FSrcEntryId
		t.FSrcBillType = inT.FSrcBillType
		t.FSrcBillNo = inT.FSrcBillNo
		t.FSrcInterId = inT.FSrcInterId
		t.FSrcEntrySeq = inT.FMoEntrySeq
		t.FMOMAINENTRYID = inT.FSrcEntryId
	}
	Q.Data.Model.FEntity = append(Q.Data.Model.FEntity, t)
}

func (Q *scrkModelBase) AddModelFEntities(ts interface{}, orgNumber string) {
	ins, ok := ts.([]*scrkEntityMini)
	if !ok {
		return
	}
	for _, inT := range ins {
		Q.addModelFEntity(inT, orgNumber)
	}
}
