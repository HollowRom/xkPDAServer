package dbTools

import (
	"fmt"
)

type SCTLMain struct {
	FBILLNO           string
	FParentNumber     string
	FParentName       string
	FParentUnitNumber string
	FUseOrgNumber     string
}

func (*SCTLMain) TableName() string {
	return "xkPdaServer_prd_ppbom_to_stockOut_tool"
}

type SCTLEntry struct {
	FID               int `json:",,string"`
	FBILLNO           string
	FParentNumber     string
	FParentName       string
	FParentUnitNumber string
	FENTRYID          int `json:",,string"`
	FSEQ              int `json:",,string"`
	FMOBILLNO         string
	FMOENTRYID        int `json:",,string"`
	FMOID             int `json:",,string"`
	FMOENTRYSEQ       int `json:",,string"`
	FMATERIALID       int `json:",,string"`
	FNUMBER           string
	FName             string
	FSPECIFICATION    string
	FBaseUnitNumber   string
	FMustQty          string
	SQTY              string
	FLOT_TEXT         string
	FUseOrgNumber     string
	FPrice            string              `xorm:"-"`
	FStockNumber      string              `xorm:"-"`
	FNote             string              `xorm:"-"`
	FStockStatusId    string              `xorm:"-"`
	FSrcBillType      string              `xorm:"-"`
	FLinkInfo         []map[string]string `xorm:"-" json:"-"`
	FISBATCHMANAGE    string
}

func (*SCTLEntry) TableName() string {
	return "xkPdaServer_prd_ppbom_to_stockOut_tool"
}

func GetAllSCTLMain(orgNumber string) []*SCTLMain {
	return getSCTLMain(orgNumber, "")
}

func GetSCTLMain(orgNumber, fBillNo string) []*SCTLMain {
	return getSCTLMain(orgNumber, fBillNo)
}

func getSCTLMain(orgNumber, fBillNo string) (r []*SCTLMain) {
	siss := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNumber))
	if fBillNo != "" {
		siss = siss.And(fmt.Sprintf("FBILLNO like '%s%%'", fBillNo))
	}
	e := siss.GroupBy("FBILLNO, FParentNumber, FParentName, FParentUnitNumber, FUseOrgNumber").Find(&r)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	if len(r) == 0 {
		fmt.Println("返回nil")
		return nil
	}

	return r
}

func GetSCTLEntry(FBillNo string) []*SCTLEntry {
	return getSCTLEntry(FBillNo)
}

func getSCTLEntry(FBillNo string) (r []*SCTLEntry) {
	e := db.Where(fmt.Sprintf("FBILLNO = '%s'", FBillNo)).Find(&r)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	if len(r) == 0 {
		fmt.Println("返回nil")
		return nil
	}

	return r
}
