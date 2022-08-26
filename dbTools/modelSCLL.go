package dbTools

import (
	"encoding/json"
	"fmt"
	"strings"
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
		FNumber string `json:"FNUMBER"`
	} `json:"FStockOrgId"`
	FPrdOrgId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FPrdOrgId"`
	FIsOwnerTInclOrg string              `json:"FIsOwnerTInclOrg"`
	FEntity          []*scllModelsEntity `json:"FInStockEntry"`
}

type scllModelsEntity struct {
	FEntryID          int `json:"FEntryID"`
	FParentMaterialId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FParentMaterialId"`
	FBaseStockActualQty string `json:"FBaseStockActualQty"`
	FMaterialId         struct {
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
	FStockStatusId struct {
		Id string `json:"Id"`
	} `json:"FStockStatusId"`
	FMoBillNo     string `json:"FMOBILLNO"`
	FMoEntryId    int    `json:"FMOENTRYID"`
	FPPBomEntryId int    `json:"FPPBomEntryId"`
	FOwnerTypeId  string `json:"FOwnerTypeId"`
	FStockUnitId  struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockUnitId"`
	FMoId        int    `json:"FMOID"`
	FMoEntrySeq  int    `json:"FMOENTRYSEQ"`
	FPPBomBillNo string `json:"FBILLNO"`
	FBaseUnitId  struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FBaseUnitId"`
	FBaseAppQty       string `json:"FBaseAppQty"`
	FBaseActualQty    string `json:"FBaseActualQty"`
	FBaseAllowOverQty string `json:"FBaseAllowOverQty"`
	FKeeperTypeId     string `json:"FKeeperTypeId"`
	FLot struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FLOT_TEXT"`
	FKeeperId         struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FKeeperId"`
	FOwnerId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FOwnerId"`
	FSrcBillType       string `json:"FSrcBillType"`
	FSrcBillNo         string `json:"FSrcBillNo"`
	FEntrySrcInterId   int    `json:"FID"`
	FEntrySrcEnteryId  int    `json:"FEntrySrcEnteryId"`
	FEntrySrcEntrySeq  int    `json:"FSEQ"`
	FParentOwnerTypeId string `json:"FParentOwnerTypeId"`
	FParentOwnerId     struct {
		FNumber string `json:"FNUMBER"`
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
	EntityMini []*SCTLEntry
	HeadMini   *SCTLMain
}

//type scllEntityMini struct {
//	FParentNumber string
//	FNUMBER         string
//	FBaseUnitNumber string
//	FMustQty        string
//	FStockNumber    string
//	FStockStatusId string
//	FMOBILLNO     string
//	FMOENTRYID    int
//	//FPPBomEntryId int
//	FMOID        int
//	FMOENTRYSEQ      int
//	FBILLNO           string
//	FID       int
//	FSEQ      int
//	FKeeperId string
//	FLOT_TEXT string
//	FLinkInfo []map[string]string
//}

//type scllHeadMini struct {
//	FUseOrgNumber string
//}

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
	inT, ok := in.(*SCTLMain)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FPrdOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FIsOwnerTInclOrg = "false"
}

func (Q *scllModelBase) addModelFEntity(inT *SCTLEntry) {
	t := &scllModelsEntity{
		FEntryID: 0,
		FParentMaterialId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FParentNumber}),
		FBaseStockActualQty: inT.FMustQty,
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
		FStockStatusId: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FMoBillNo:     inT.FMOBILLNO,
		FMoEntryId:    inT.FMOENTRYID,
		//FPPBomEntryId: inT.FPPBomEntryId,
		FOwnerTypeId:  "BD_OwnerOrg",
		FStockUnitId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FMoId:        inT.FMOID,
		FMoEntrySeq:  inT.FMOENTRYSEQ,
		FLot: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: strings.TrimRight(inT.FLOT_TEXT, " ")}),
		FPPBomBillNo: inT.FBILLNO,
		FBaseUnitId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FBaseAppQty:       inT.FMustQty,
		FBaseActualQty:    inT.FMustQty,
		FBaseAllowOverQty: inT.FMustQty,
		FKeeperTypeId:     "BD_KeeperOrg",
		FKeeperId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FUseOrgNumber}),
		FOwnerId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FUseOrgNumber}),
		FSrcBillType:       "PRD_PPBOM",
		FSrcBillNo:         inT.FBILLNO,
		FEntrySrcInterId:   inT.FID,
		//FEntrySrcEnteryId:  inT.FPPBomEntryId,
		FEntrySrcEntrySeq:  inT.FSEQ,
		FParentOwnerTypeId: "BD_OwnerOrg",
		FParentOwnerId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FUseOrgNumber}),
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

func (Q *scllModelBase) AddModelFEntities(ts interface{}) {
	in, ok := ts.([]*SCTLEntry)
	if !ok {
		return
	}
	for _, inT := range in {
		Q.addModelFEntity(inT)
	}
}
