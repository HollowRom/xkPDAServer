package dbTools

import "time"

type ModelBaseInterface interface {
	CheckVerify() bool
	GetJson() []byte
	AddModelFEntities(interface{})
	AddModelHead(interface{})
}

type DefModelHeadBase struct {
	FBillTypeId string
	FDate       time.Time
	FromId      string
}

type DefDataHeadBase struct {
	NeedUpDateFields      []string
	NeedReturnFields      []string
	IsDeleteEntry         string
	SubSystemId           string
	IsVerifyBaseDataField string
	IsEntryBatchFill      string
	ValidateFlag          string
	NumberSearch          string
	IsAutoAdjustField     string
	InterationFlags       string
	IgnoreInterationFlag  string
}
