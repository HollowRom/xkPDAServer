package dbTools

import (
	"encoding/json"
	"fmt"
	"strings"
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
		FNumber string `json:"FNUMBER"`
	} `json:"FStockOrgId"`
	FPrdOrgId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FPrdOrgId"`
	FOwnerTypeId string `json:"FOwnerTypeId"`
	FOwnerId0    struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FOwnerId0"`
	FEntity []*scrkModelsEntity `json:"FEntity"`
}

type scrkModelsEntity struct {
	FSrcEntryId int `json:"FSrcEntryId"`
	FMaterialId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FMaterialId"`
	FUnitID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FUnitID"`
	FMustQty    string `json:"FMustQty"`
	FRealQty    string `json:"FRealQty"`
	FBaseUnitId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FBaseUnitId"`
	FOwnerTypeId string `json:"FOwnerTypeId"`
	FOwnerId     struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FOwnerId"`
	FStockId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockId"`
	FLot struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FLOT_TEXT"`
	FMoBillNo    string `json:"FMOBILLNO"`
	FMoId        int    `json:"FMOID"`
	FMoEntryId   int    `json:"FMOENTRYID"`
	FMoEntrySeq  int    `json:"FMOENTRYSEQ"`
	FStockUnitId struct {
		FNumber string `json:"FNUMBER"`
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
		FNumber string `json:"FNUMBER"`
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
	EntityMini []*SCDDEntry `json:"FEntity"`
	//HeadMini   *SCDDMain `json:"HeadMini"`
}

//type scrkEntityMini struct {
//	FNUMBER    string
//	FENTRYID        int
//	FBaseUnitNumber string
//	FMustQty        string
//	FPrice          string
//	FStockNumber   string
//	FNote          string
//	FStockStatusId string
//	FLOT_TEXT string
//	FMOBILLNO string
//	FMOID       int
//	FMOENTRYID  int
//	FMOENTRYSEQ  int
//	FBILLNO      string
//	FSrcBillType string
//	FID          int
//	FLinkInfo    []map[string]string
//}

//type scrkHeadMini struct {
//	FUseOrgNumber string
//}

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
	inT, ok := in.(*SCDDEntry)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FPrdOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FPrdOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FOwnerTypeId = "BD_OwnerOrg"
	Q.Data.Model.FOwnerId0.FNumber = inT.FUseOrgNumber
}

func (Q *scrkModelBase) addModelFEntity(inT *SCDDEntry) {
	t := &scrkModelsEntity{
		FSrcEntryId: inT.FENTRYID,
		FMaterialId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FNUMBER}),
		FUnitID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FMustQty: inT.SQTY,
		FBaseUnitId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FRealQty:     inT.SQTY,
		FOwnerTypeId: "BD_OwnerOrg",
		FOwnerId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FUseOrgNumber}),
		FStockId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FLot: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: strings.TrimRight(inT.FLOT_TEXT, " ")}),
		FStockUnitId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FStockStatusId: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FKeeperTypeId: "BD_KeeperOrg",
		FKeeperId: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FUseOrgNumber}),
	}

	t.FentityLink = append(t.FentityLink, &scrkFENTITY_Link{
		FENTITYLinkFRuleId:        inT.FLinkInfo[0]["FENTITY_Link_FRuleId"],
		FENTITYLinkFSTableName:    inT.FLinkInfo[0]["FENTITY_Link_FSTableName"],
		FENTITYLinkFSBillId:       inT.FLinkInfo[0]["FENTITY_Link_FSBillId"],
		FENTITYLinkFSId:           inT.FLinkInfo[0]["FENTITY_Link_FSId"],
		FENTITYLinkFBaseActualQty: inT.FLinkInfo[0]["FENTITY_Link_FBaseActualQty"],
	})
	t.FMoBillNo = inT.FBILLNO
	t.FMoId = inT.FID
	t.FMoEntrySeq = inT.FSEQ
	t.FSrcEntryId = inT.FENTRYID
	t.FMoEntryId = inT.FENTRYID
	t.FSrcBillType = inT.FSrcBillType
	t.FSrcBillNo = inT.FBILLNO
	t.FSrcInterId = inT.FID
	t.FSrcEntrySeq = inT.FSEQ
	t.FMOMAINENTRYID = inT.FENTRYID

	Q.Data.Model.FEntity = append(Q.Data.Model.FEntity, t)
}

func (Q *scrkModelBase) AddModelFEntities(ts interface{}) {
	ins, ok := ts.([]*SCDDEntry)
	if !ok {
		return
	}
	for _, inT := range ins {
		Q.addModelFEntity(inT)
	}
}
