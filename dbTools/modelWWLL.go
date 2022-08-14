package dbTools

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
		FNumber string `json:"FNUMBER"`
	} `json:"FStockOrgId"`
	FSubOrgId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSubOrgId"`
	FEntity []*wwllModelsEntity `json:"FInStockEntry"`
}

type wwllModelsEntity struct {
	FMaterialId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FMaterialId"`
	FUnitID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FUnitID"`
	FAppQty    string `json:"FAppQty"`
	FActualQty string `json:"FActualQty"`
	FStockId   struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockId"`
	FParentMaterialId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FParentMaterialId"`
	FSrcInterId     int    `json:"FID"`
	FSrcEntryId     int    `json:"FENTRYID"`
	FSrcBillType    string `json:"FSrcBillType"`
	FSrcBillNo      string `json:"FBILLNO"`
	FSrcEntrySeq    int    `json:"FSrcEntrySeq"`
	FSubReqId       int    `json:"FSUBREQID"`
	FPPbomBillNo    string `json:"FBILLNO"`
	FSubReqEntryId  int    `json:"FSubReqEntryId"`
	FSubReqBillNo   string `json:"FSUBREQBILLNO"`
	FSubReqEntrySeq int    `json:"FSubReqEntrySeq"`
	FBaseUnitId     struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FBaseUnitId"`
	FBaseAppQty    string `json:"FBaseAppQty"`
	FBaseActualQty string `json:"FBaseActualQty"`
	FStockUnitId   struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockUnitId"`
	FStockAppQty           string `json:"FStockAppQty"`
	FStockAllowOverQty     string `json:"FStockAllowOverQty"`
	FStockSelPrcdReturnQty string `json:"FStockSelPrcdReturnQty"`
	FPPbomEntryId          int    `json:"FENTRYID"`
	FPOOrderBillNo         string `json:"FBILLNO"`
	FPOOrderSeq            int    `json:"FPOOrderSeq"`
	FSupplierId            struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSupplierId"`
	FKeeperTypeId string `json:"FKeeperTypeId"`
	FKeeperId     struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FKeeperId"`
	FStockStatusId struct {
		Id string `json:"Id"`
	} `json:"FStockStatusId"`
	FOwnerTypeId string `json:"FOwnerTypeId"`
	FOwnerId     struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FOwnerId"`
	FParentOwnerTypeId string `json:"FParentOwnerTypeId"`
	FParentOwnerId     struct {
		FNumber string `json:"FNUMBER"`
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
	WWLLEntityMini []*WWTLEntry
	WWLLHeadMini   *WWTLMain
}

//type wwllEntityMini struct {
//	FParentNumber   string
//	FNUMBER         string
//	FBaseUnitNumber string
//	FMustQty        string
//	FStockNumber    string
//	FStockStatusId  string
//	FSrcInterId     int
//	FSrcEntryId     int
//	FSrcBillNo      string
//	FSrcEntrySeq    int
//	FSUBREQID       int
//	FBILLNO         string
//	FSUBREQBILLNO   string
//	FENTRYID        int
//	FPOOrderBillNo  string
//	FKeeperId       string
//	FLOT_TEXT       string
//	FLinkInfo       []map[string]string
//}

//type wwllHeadMini struct {
//	FBILLNO       string
//	FUseOrgNumber string
//	FSuppNumber   string
//	FSuppName   string
//}

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
	inT, ok := in.(*WWTLMain)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FBillType.Id = inT.FUseOrgNumber
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FSubOrgId.FNumber = inT.FUseOrgNumber
}

func (Q *wwllModelBase) addModelFEntity(inT *WWTLEntry, orgNumber string) {
	t := &wwllModelsEntity{
		FParentMaterialId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FParentNumber}),
		FUnitID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FMaterialId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FNUMBER}),
		FAppQty:    inT.FMustQty,
		FActualQty: inT.FMustQty,
		FStockId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FSrcInterId:    inT.FSrcInterId,
		FSrcEntryId:    inT.FSrcEntryId,
		FSrcBillType:   "SUB_PPBOM",
		FSrcBillNo:     inT.FSrcBillNo,
		FSrcEntrySeq:   inT.FSrcEntrySeq,
		FSubReqId:      inT.FSUBREQID,
		FPPbomBillNo:   inT.FBILLNO,
		FSubReqEntryId: inT.FSUBREQID,
		FSubReqBillNo:  inT.FSUBREQBILLNO,
		FBaseUnitId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FBaseAppQty:    inT.FMustQty,
		FBaseActualQty: inT.FMustQty,
		FStockUnitId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FStockAppQty:           inT.FMustQty,
		FStockAllowOverQty:     inT.FMustQty,
		FStockSelPrcdReturnQty: inT.FMustQty,
		FPPbomEntryId:          inT.FENTRYID,
		FPOOrderBillNo:         inT.FPOOrderBillNo,
		FKeeperTypeId:          "BD_KeeperOrg",
		FKeeperId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
		FStockStatusId: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FOwnerTypeId: "BD_OwnerOrg",
		FOwnerId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
		FParentOwnerTypeId: "BD_OwnerOrg",
		FParentOwnerId: struct {
			FNumber string `json:"FNUMBER"`
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
	in, ok := ts.([]*WWTLEntry)
	if !ok {
		return
	}
	for _, inT := range in {
		Q.addModelFEntity(inT, orgNumber)
	}
}
