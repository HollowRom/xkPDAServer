package jsonTools

import (
	"encoding/json"
	"fmt"
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
		FNumber string `json:"FNumber"`
	} `json:"FStockOrgId"`
	FStockDirect string `json:"FStockDirect"`
	FDate        string `json:"FDate"`
	FSUPPLIERID  struct {
		FNumber string `json:"FNumber"`
	} `json:"FSUPPLIERID"`
	FEntity []*qtrkModelsEntity `json:"FInStockEntry"`
}

type qtrkModelsEntity struct {
	FMATERIALID struct {
		FNumber string `json:"FNumber"`
	} `json:"FMATERIALID"`
	FUnitID struct {
		FNumber string `json:"FNumber"`
	} `json:"FUnitID"`
	FSTOCKID struct {
		FNumber string `json:"FNumber"`
	} `json:"FSTOCKID"`
	FSTOCKSTATUSID struct {
		Id string `json:"Id"`
	} `json:"FSTOCKSTATUSID"`
	FQty         string `json:"FQty"`
	FOWNERTYPEID string `json:"FOWNERTYPEID"`
	FOWNERID     struct {
		FNumber string `json:"FNumber"`
	} `json:"FOWNERID"`
}

type QTRKMini struct {
	QTRKEntityMini []*QTRKEntityMini
	QTRKHeadMini   *QTRKHeadMini
}

type QTRKEntityMini struct {
	FNumber        string
	UnitNumber     string
	FQTY           string
	FStockStatusId string
	FLotNo         string
	FLinkInfo      []map[string]string
}

type QTRKHeadMini struct {
	FOrgNumber   string
	FStockDirect string
	FSUPPLIERID  string
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
	inT, ok := in.(*QTRKHeadMini)
	if !ok {
		return
	}
	Q.Data.Model.FStockOrgId.FNumber = inT.FOrgNumber
	Q.Data.Model.FSUPPLIERID.FNumber = inT.FSUPPLIERID
}

func (Q *qtrkModelBase) addModelFEntity(inT *QTRKEntityMini, orgNumber string) {
	t := &qtrkModelsEntity{
		FMATERIALID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.FNumber}),
		FUnitID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: inT.UnitNumber}),
		FQty:         inT.FQTY,
		FOWNERTYPEID: "BD_OwnerOrg",
		FSTOCKSTATUSID: struct {
			Id string `json:"Id"`
		}(struct{ Id string }{Id: inT.FStockStatusId}),
		FOWNERID: struct {
			FNumber string `json:"FNumber"`
		}(struct{ FNumber string }{FNumber: orgNumber}),
	}
	Q.Data.Model.FEntity = append(Q.Data.Model.FEntity, t)
}

func (Q *qtrkModelBase) AddModelFEntities(ts interface{}, orgNumber string) {
	in, ok := ts.([]*QTRKEntityMini)
	if !ok {
		return
	}
	for _, inT := range in {
		Q.addModelFEntity(inT, orgNumber)
	}
}
