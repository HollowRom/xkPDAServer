package jsonTools

import (
	"encoding/json"
	"fmt"
)

type wwllModelBase struct {
	Formid string         `json:"formid"`
	Data   *wwllModelData `json:"data"`
}

type wwllModelData struct {
	Model *wwllModels `json:"Model"`
}

type wwllModels struct {
	FBillType struct {
		Id string `json:"Id"`
	} `json:"FBillType"`
	FDate       string `json:"FDate"`
	FStockOrgId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockOrgId"`
	FSubOrgId struct {
		FNumber string `json:"FNumber"`
	} `json:"FSubOrgId"`
	FEntity []*wwllModelsEntity `json:"FInStockEntry"`
}

type wwllModelsEntity struct {
	FMaterialId struct {
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
	FParentMaterialId struct {
		FNumber string `json:"FNumber"`
	} `json:"FParentMaterialId"`
	FSrcInterId     int    `json:"FSrcInterId"`
	FSrcEntryId     int    `json:"FSrcEntryId"`
	FSrcBillType    string `json:"FSrcBillType"`
	FSrcBillNo      string `json:"FSrcBillNo"`
	FSrcEntrySeq    int    `json:"FSrcEntrySeq"`
	FSubReqId       int    `json:"FSubReqId"`
	FPPbomBillNo    string `json:"FPPbomBillNo"`
	FSubReqEntryId  int    `json:"FSubReqEntryId"`
	FSubReqBillNo   string `json:"FSubReqBillNo"`
	FSubReqEntrySeq int    `json:"FSubReqEntrySeq"`
	FBaseUnitId     struct {
		FNumber string `json:"FNumber"`
	} `json:"FBaseUnitId"`
	FBaseAppQty    string `json:"FBaseAppQty"`
	FBaseActualQty string `json:"FBaseActualQty"`
	FStockUnitId   struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockUnitId"`
	FStockAppQty           string `json:"FStockAppQty"`
	FStockAllowOverQty     string `json:"FStockAllowOverQty"`
	FStockSelPrcdReturnQty string `json:"FStockSelPrcdReturnQty"`
	FPPbomEntryId          int    `json:"FPPbomEntryId"`
	FPOOrderBillNo         string `json:"FPOOrderBillNo"`
	FPOOrderSeq            int    `json:"FPOOrderSeq"`
	FSupplierId            struct {
		FNumber string `json:"FNumber"`
	} `json:"FSupplierId"`
	FKeeperTypeId string `json:"FKeeperTypeId"`
	FKeeperId     struct {
		FNumber string `json:"FNumber"`
	} `json:"FKeeperId"`
	FStockStatusId struct {
		Id string `json:"Id"`
	} `json:"FStockStatusId"`
	FOwnerTypeId string `json:"FOwnerTypeId"`
	FOwnerId     struct {
		FNumber string `json:"FNumber"`
	} `json:"FOwnerId"`
	FParentOwnerTypeId string `json:"FParentOwnerTypeId"`
	FParentOwnerId     struct {
		FNumber string `json:"FNumber"`
	} `json:"FParentOwnerId"`
	FentityLink []*wwllFEntity_Link `json:"FEntity_Link"`
}

type wwllFEntity_Link struct {
	FEntityLinkFRuleId        string `json:"FEntity_Link_FRuleId"`
	FEntityLinkFSTableName    string `json:"FEntity_Link_FSTableName"`
	FEntityLinkFSBillId       string `json:"FEntity_Link_FSBillId"`
	FEntityLinkFSId           string `json:"FEntity_Link_FSId"`
	FEntityLinkFBaseActualQty string `json:"FEntity_Link_FBaseActualQty"`
}

type WWLLMini struct {
	WWLLEntityMini []*wwllEntityMini
	WWLLHeadMini   *wwllHeadMini
}

type wwllEntityMini struct {
	FParentNumber  string
	FNumber        string
	UnitNumber     string
	FQTY           string
	FStockNumber   string
	FStockStatusId string
	FSrcInterId    int
	FSrcEntryId    int
	FSrcBillNo     string
	FSrcEntrySeq   int
	FSubReqId      int
	FPPbomBillNo   string
	FSubReqBillNo  string
	FPPbomEntryId  int
	FPOOrderBillNo string
	FKeeperId      string
	FLotNo         string
	FLinkInfo      []map[string]string
}

type wwllHeadMini struct {
	FBillNo         string
	FOrgNumber      string
	FSUPPLIERNumber string
	FSUPPLIERName   string
}

var _ ModelBaseInterface = &wwllModelBase{}

func InitwwllModel(initBase *DefModelHeadBase) *wwllModelBase {
	if initBase == nil {
		return nil
	}
	return &wwllModelBase{Formid: initBase.FromId, Data: &wwllModelData{
		Model: &wwllModels{
			FBillType: struct {
				Id string `json:"Id"`
			}{Id: initBase.FBillTypeId},
			FDate: initBase.FDate.Format("2006-01-02"),
		},
	}}
}

func (Q *wwllModelBase) CheckVerify() bool {
	return true
}

func (Q *wwllModelBase) GetJson() []byte {
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

func (Q *wwllModelBase) AddModelHead(in interface{}) {
	inT, ok := in.(*wwllHeadMini)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FBillType.Id = inT.FOrgNumber
	Q.Data.Model.FStockOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FSubOrgId.FNumber = inT.FOrgNumber
}

func (Q *wwllModelBase) addModelFEntity(inT *wwllEntityMini, orgNumber string) {
	t := &wwllModelsEntity{
		FParentMaterialId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FParentNumber}),
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
		FSrcInterId:    inT.FSrcInterId,
		FSrcEntryId:    inT.FSrcEntryId,
		FSrcBillType:   "SUB_PPBOM",
		FSrcBillNo:     inT.FSrcBillNo,
		FSrcEntrySeq:   inT.FSrcEntrySeq,
		FSubReqId:      inT.FSubReqId,
		FPPbomBillNo:   inT.FPPbomBillNo,
		FSubReqEntryId: inT.FSubReqId,
		FSubReqBillNo:  inT.FSubReqBillNo,
		FBaseUnitId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FBaseAppQty:    inT.FQTY,
		FBaseActualQty: inT.FQTY,
		FStockUnitId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FStockAppQty:           inT.FQTY,
		FStockAllowOverQty:     inT.FQTY,
		FStockSelPrcdReturnQty: inT.FQTY,
		FPPbomEntryId:          inT.FPPbomEntryId,
		FPOOrderBillNo:         inT.FPOOrderBillNo,
		FKeeperTypeId:          "BD_KeeperOrg",
		FKeeperId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
		FStockStatusId: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FOwnerTypeId: "BD_OwnerOrg",
		FOwnerId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
		FParentOwnerTypeId: "BD_OwnerOrg",
		FParentOwnerId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
	}
	t.FentityLink = append(t.FentityLink, &wwllFEntity_Link{
		FEntityLinkFRuleId:        inT.FLinkInfo[0]["FEntity_Link_FRuleId"],
		FEntityLinkFSTableName:    inT.FLinkInfo[0]["FEntity_Link_FSTableName"],
		FEntityLinkFSBillId:       inT.FLinkInfo[0]["FEntity_Link_FSBillId"],
		FEntityLinkFSId:           inT.FLinkInfo[0]["FEntity_Link_FSId"],
		FEntityLinkFBaseActualQty: inT.FLinkInfo[0]["FEntity_Link_FBaseActualQty"],
	})
	Q.Data.Model.FEntity = append(Q.Data.Model.FEntity, t)
}

func (Q *wwllModelBase) AddModelFEntities(ts interface{}, orgNumber string) {
	in, ok := ts.([]*wwllEntityMini)
	if !ok {
		return
	}
	for _, inT := range in {
		Q.addModelFEntity(inT, orgNumber)
	}
}
