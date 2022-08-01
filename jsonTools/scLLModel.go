package jsonTools

import (
	"encoding/json"
	"fmt"
)

type scllModelBase struct {
	Formid string         `json:"formid"`
	Data   *scllModelData `json:"data"`
}

type scllModelData struct {
	Model *scllModels `json:"Model"`
}

type scllModels struct {
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
	FIsOwnerTInclOrg string              `json:"FIsOwnerTInclOrg"`
	FEntity          []*scllModelsEntity `json:"FInStockEntry"`
}

type scllModelsEntity struct {
	FEntryID          int `json:"FEntryID"`
	FParentMaterialId struct {
		FNumber string `json:"FNumber"`
	} `json:"FParentMaterialId"`
	FBaseStockActualQty string `json:"FBaseStockActualQty"`
	FMaterialId         struct {
		FNumber string `json:"FNumber"`
	} `json:"FMaterialId"`
	FUnitID struct {
		FNumber string `json:"FNumber"`
	} `json:"FUnitID"`
	FAppQty    string `json:"FAppQty"`
	FActualQty string `json:"FActualQty"`
	FStockId   struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockId"`
	FStockStatusId struct {
		Id string `json:"Id"`
	} `json:"FStockStatusId"`
	FMoBillNo     string `json:"FMoBillNo"`
	FMoEntryId    int    `json:"FMoEntryId"`
	FPPBomEntryId int    `json:"FPPBomEntryId"`
	FOwnerTypeId  string `json:"FOwnerTypeId"`
	FStockUnitId  struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockUnitId"`
	FMoId        int    `json:"FMoId"`
	FMoEntrySeq  int    `json:"FMoEntrySeq"`
	FPPBomBillNo string `json:"FPPBomBillNo"`
	FBaseUnitId  struct {
		FNumber string `json:"FNumber"`
	} `json:"FBaseUnitId"`
	FBaseAppQty       string `json:"FBaseAppQty"`
	FBaseActualQty    string `json:"FBaseActualQty"`
	FBaseAllowOverQty string `json:"FBaseAllowOverQty"`
	FKeeperTypeId     string `json:"FKeeperTypeId"`
	FKeeperId         struct {
		FNumber string `json:"FNumber"`
	} `json:"FKeeperId"`
	FOwnerId struct {
		FNumber string `json:"FNumber"`
	} `json:"FOwnerId"`
	FSrcBillType       string `json:"FSrcBillType"`
	FSrcBillNo         string `json:"FSrcBillNo"`
	FEntrySrcInterId   int    `json:"FEntrySrcInterId"`
	FEntrySrcEnteryId  int    `json:"FEntrySrcEnteryId"`
	FEntrySrcEntrySeq  int    `json:"FEntrySrcEntrySeq"`
	FParentOwnerTypeId string `json:"FParentOwnerTypeId"`
	FParentOwnerId     struct {
		FNumber string `json:"FNumber"`
	} `json:"FParentOwnerId"`
	FentityLink []*scllFEntity_Link `json:"FEntity_Link"`
}

type scllFEntity_Link struct {
	FentityLinkFruleid         string `json:"FEntity_Link_FRuleId"`
	FentityLinkFstablename     string `json:"FEntity_Link_FSTableName"`
	FentityLinkFsbillid        string `json:"FEntity_Link_FSBillId"`
	FentityLinkFsid            string `json:"FEntity_Link_FSId"`
	FentityLinkFbaseprdrealqty string `json:"FEntity_Link_FBasePrdRealQty"`
}

type SCLLMini struct {
	SCLLEntityMini []*scllEntityMini
	SCLLHeadMini   *scllHeadMini
}

type scllEntityMini struct {
	FParentNumber     string
	FNumber           string
	UnitNumber        string
	FQTY              string
	FStockNumber      string
	FStockStatusId    string
	FMoBillNo         string
	FMoEntryId        int
	FPPBomEntryId     int
	FMoId             int
	FMoEntrySeq       int
	FPPBomBillNo      string
	FEntrySrcInterId  int
	FEntrySrcEntrySeq int
	FKeeperId         string
	FLotNo            string
	FLinkInfo         []map[string]string
}

type scllHeadMini struct {
	FOrgNumber string
}

var _ ModelBaseInterface = &scllModelBase{}

func InitScllModel(initBase *DefModelHeadBase) *scllModelBase {
	if initBase == nil {
		return nil
	}
	return &scllModelBase{Formid: initBase.FromId, Data: &scllModelData{
		Model: &scllModels{
			FBillType: struct {
				Id string `json:"Id"`
			}{Id: initBase.FBillTypeId},
			FDate: initBase.FDate.Format("2006-01-02"),
		},
	}}
}

func (Q *scllModelBase) CheckVerify() bool {
	return true
}

func (Q *scllModelBase) GetJson() []byte {
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

func (Q *scllModelBase) AddModelHead(in interface{}) {
	inT, ok := in.(*scllHeadMini)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FPrdOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FIsOwnerTInclOrg = "false"
}

func (Q *scllModelBase) addModelFEntity(inT *scllEntityMini, orgNumber string) {
	t := &scllModelsEntity{
		FEntryID: 0,
		FParentMaterialId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FParentNumber}),
		FBaseStockActualQty: inT.FQTY,
		FUnitID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FMaterialId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FNumber}),
		FAppQty:    inT.FQTY,
		FActualQty: inT.FQTY,
		FStockId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FStockStatusId: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FMoBillNo:     inT.FMoBillNo,
		FMoEntryId:    inT.FMoEntryId,
		FPPBomEntryId: inT.FPPBomEntryId,
		FOwnerTypeId:  "BD_OwnerOrg",
		FStockUnitId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FMoId:        inT.FMoId,
		FMoEntrySeq:  inT.FMoEntrySeq,
		FPPBomBillNo: inT.FPPBomBillNo,
		FBaseUnitId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FBaseAppQty:       inT.FQTY,
		FBaseActualQty:    inT.FQTY,
		FBaseAllowOverQty: inT.FQTY,
		FKeeperTypeId:     "BD_KeeperOrg",
		FKeeperId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
		FOwnerId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
		FSrcBillType:       "PRD_PPBOM",
		FSrcBillNo:         inT.FPPBomBillNo,
		FEntrySrcInterId:   inT.FEntrySrcInterId,
		FEntrySrcEnteryId:  inT.FPPBomEntryId,
		FEntrySrcEntrySeq:  inT.FEntrySrcEntrySeq,
		FParentOwnerTypeId: "BD_OwnerOrg",
		FParentOwnerId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
	}
	t.FentityLink = append(t.FentityLink, &scllFEntity_Link{
		FentityLinkFruleid:         inT.FLinkInfo[0]["FEntity_Link_FRuleId"],
		FentityLinkFstablename:     inT.FLinkInfo[0]["FEntity_Link_FSTableName"],
		FentityLinkFsbillid:        inT.FLinkInfo[0]["FEntity_Link_FSBillId"],
		FentityLinkFsid:            inT.FLinkInfo[0]["FEntity_Link_FSId"],
		FentityLinkFbaseprdrealqty: inT.FLinkInfo[0]["FEntity_Link_FBasePrdRealQty"],
	})
	Q.Data.Model.FEntity = append(Q.Data.Model.FEntity, t)
}

func (Q *scllModelBase) AddModelFEntities(ts interface{}, orgNumber string) {
	in, ok := ts.([]*scllEntityMini)
	if !ok {
		return
	}
	for _, inT := range in {
		Q.addModelFEntity(inT, orgNumber)
	}
}
