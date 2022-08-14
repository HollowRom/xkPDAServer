package dbTools

import (
	"fmt"
)

type XSDDMain struct {
	FBILLNO     string
	FCustNumber string
	FCustName   string
	FUseOrgNumber string
}

func (*XSDDMain) TableName() string {
	return "xkPdaServer_sale_tz_to_stockout_tool"
}

type XSDDEntry struct {
	FID            int `json:",,string"`
	FBILLNO        string
	FCustNumber    string
	FCustName      string
	FENTRYID       int `json:",,string"`
	FSEQ           int `json:",,string"`
	FORDERNO       string
	FORDERSEQ      int `json:",,string"`
	FMATERIALID    int `json:",,string"`
	FNUMBER        string
	FNAME          string
	FBaseUnitNumber string
	FSPECIFICATION string
	FLOT_TEXT      string
	FMustQty       string
	SQTY           string
	FUseOrgNumber  string
	FLinkInfo      []map[string]string `xorm:"-" json:"-"`
	FStockNumber string `xorm:"-"`
	FPrice string `xorm:"-"`
	FStockStatusId string `xorm:"-"`
}

func (*XSDDEntry) TableName() string {
	return "xkPdaServer_sale_tz_to_stockout_tool"
}

func GetAllXSDDMain(orgNumber string) []*XSDDMain {
	return getXSDDMain(orgNumber, "", "")
}

func GetXSDDMain(orgNumber, custNumber, fBillNo string) []*XSDDMain {
	return getXSDDMain(orgNumber, custNumber, fBillNo)
}

func getXSDDMain(orgNumber, custNumber, fBillNo string) (r []*XSDDMain) {
	siss := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNumber))
	if custNumber != "" {
		siss = siss.And(fmt.Sprintf("FCustNumber = '%s'", custNumber))
	}
	if fBillNo != "" {
		siss = siss.And(fmt.Sprintf("FBILLNO like '%s%%'", fBillNo))
	}
	e := siss.GroupBy("FBILLNO, FCustNumber, FCustName, FUseOrgNumber").Find(&r)
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

func GetXSDDEntry(FBillNo string) []*XSDDEntry {
	return getXSDDEntry(FBillNo)
}

func getXSDDEntry(FBillNo string) (r []*XSDDEntry) {
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
