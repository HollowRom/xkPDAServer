package dbTools

import (
	"encoding/json"
	"fmt"
	"strings"
)

type qtckModelBase struct {
	Formid string         `json:"formid"`
	Data   *qtckModelData `json:"data"`
}

type qtckModelData struct {
	Model *qtckModels `json:"Model"`
}

type qtckModels struct {
	FBillTypeID struct {
		Id string `json:"Id"`
	} `json:"FBillTypeID"`
	FStockOrgId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockOrgId"`
	FStockDirect string `json:"FStockDirect"`
	FDate        string `json:"FDate"`
	FCustId      struct {
		FNumber string `json:"FNumber"`
	} `json:"FCustId"`
	FEntity []*qtckModelsEntity `json:"FEntity"`
}

type qtckModelsEntity struct {
	FMaterialId struct {
		FNumber string `json:"FNumber"`
	} `json:"FMaterialId"`
	FUnitID struct {
		FNumber string `json:"FNumber"`
	} `json:"FUnitID"`
	FQty        string `json:"FQty"`
	FBaseUnitId struct {
		FNumber string `json:"FNumber"`
	} `json:"FBaseUnitId"`
	FStockId struct {
		FNumber string `json:"FNumber"`
	} `json:"FStockId"`
	FLot struct {
		FNumber string `json:"FNumber"`
	} `json:"FLot"`
	FOwnerTypeId string `json:"FOwnerTypeId"`
	FOwnerId     struct {
		FNumber string `json:"FNumber"`
	} `json:"FOwnerId"`
	FStockStatusId struct {
		Id string `json:"Id"`
	} `json:"FStockStatusId"`
}

type QTCKMini struct {
	EntityMini []*QTCRKEntry
}

var _ ModelBaseInterface = &qtckModelBase{}

func InitQTCKModel(initBase *DefModelHeadBase) *qtckModelBase {
	if initBase == nil {
		return nil
	}
	return &qtckModelBase{Formid: initBase.FromId, Data: &qtckModelData{
		Model: &qtckModels{
			FBillTypeID: struct {
				Id string `json:"Id"`
			}(struct{ Id string }{Id: initBase.FBillTypeId}),
			FStockDirect: "GENERAL",
			FDate:        initBase.FDate.Format("2006-01-02"),
		},
	}}
}

func (Q *qtckModelBase) CheckVerify() bool {
	return true
}

func (Q *qtckModelBase) GetJson() []byte {
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

func (Q *qtckModelBase) AddModelHead(in interface{}) {
	inT, ok := in.(*QTCRKEntry)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FUseOrgNumber
	Q.Data.Model.FCustId.FNumber = inT.FCustNumber
}

func (Q *qtckModelBase) addModelFEntity(inT *QTCRKEntry) {
	t := &qtckModelsEntity{
		FMaterialId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FNumber}),
		FUnitID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FQty: inT.FQTY,
		FBaseUnitId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FStockId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
		FLot: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: strings.TrimRight(inT.FLOT_TEXT, " ")}),
		FOwnerTypeId: "BD_OwnerOrg",
		FOwnerId: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FUseOrgNumber}),
		FStockStatusId: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
	}
	Q.Data.Model.FEntity = append(Q.Data.Model.FEntity, t)
}

func (Q *qtckModelBase) AddModelFEntities(ts interface{}) {
	ins, ok := ts.([]*QTCRKEntry)
	if !ok {
		return
	}
	for _, inT := range ins {
		Q.addModelFEntity(inT)
	}
}
