package dbTools

import (
	"encoding/json"
	"fmt"
	"strings"
)

type qtrkModelBase struct {
	Formid string         `json:"formid"`
	Data   *qtrkModelData `json:"data"`
}

type qtrkModelData struct {
	Model *qtrkModels `json:"Model"`
}

type qtrkModels struct {
	FBillTypeID struct {
		Id string `json:"Id"`
	} `json:"FBillTypeID"`
	FStockOrgId struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FStockOrgId"`
	FStockDirect string `json:"FStockDirect"`
	FDate        string `json:"FDate"`
	FSUPPLIERID  struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSUPPLIERID"` //合并了出入库单所以供应商和客户分不开
	FEntity []*qtrkModelsEntity `json:"FEntity"`
}

type qtrkModelsEntity struct {
	FMATERIALID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FMATERIALID"`
	FUnitID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FUnitID"`
	FSTOCKID struct {
		FNumber string `json:"FNUMBER"`
	} `json:"FSTOCKID"`
	FSTOCKSTATUSID struct {
		Id string `json:"Id"`
	} `json:"FSTOCKSTATUSID"`
	FLOT struct {
		FNumber string `json:"FNumber"`
	} `json:"FLOT"`
	FQty         string `json:"FQty"`
	FOWNERTYPEID string `json:"FOWNERTYPEID"`
	FOWNERID     struct {
		Id string `json:"Id"`
	} `json:"FOWNERID"`
}

type QTRKMini struct {
	EntityMini []*QTCRKEntry
}

var _ ModelBaseInterface = &qtrkModelBase{}

func InitQTRKModel(initBase *DefModelHeadBase) *qtrkModelBase {
	if initBase == nil {
		return nil
	}
	return &qtrkModelBase{Formid: initBase.FromId, Data: &qtrkModelData{
		Model: &qtrkModels{
			FBillTypeID: struct {
				Id string `json:"Id"`
			}{Id: initBase.FBillTypeId},
			FStockDirect: "GENERAL",
			FDate:        initBase.FDate.Format("2006-01-02"),
		},
	}}
}

func (Q *qtrkModelBase) CheckVerify() bool {
	return true
}

func (Q *qtrkModelBase) GetJson() []byte {
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

func (Q *qtrkModelBase) AddModelHead(in interface{}) {
	inT, ok := in.([]*QTCRKEntry)
	if !ok || len(inT) < 1 {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT[0].FUseOrgNumber
	Q.Data.Model.FSUPPLIERID.FNumber = inT[0].FCustNumber
}

func (Q *qtrkModelBase) addModelFEntity(inT *QTCRKEntry) {
	t := &qtrkModelsEntity{
		FMATERIALID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FNumber}),
		FUnitID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FBaseUnitNumber}),
		FQty:         inT.FQTY,
		FOWNERTYPEID: "BD_OwnerOrg",
		FSTOCKSTATUSID: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FLOT: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: strings.TrimRight(inT.FLOT_TEXT, " ")}),
		FOWNERID: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FUseOrgNumber}),
		FSTOCKID: struct {
			FNumber string `json:"FNUMBER"`
		}(struct{ FNumber string }{FNumber: inT.FStockNumber}),
	}
	Q.Data.Model.FEntity = append(Q.Data.Model.FEntity, t)
}

func (Q *qtrkModelBase) AddModelFEntities(ts interface{}) {
	in, ok := ts.([]*QTCRKEntry)
	if !ok {
		return
	}
	for _, inT := range in {
		Q.addModelFEntity(inT)
	}
}
