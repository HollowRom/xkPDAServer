package dbTools

import (
	"fmt"
)

type SCDDMain struct {
	FBILLNO       string
	FUseOrgNumber string
	FNUMBER string
}

func (*SCDDMain) TableName() string {
	return "xkPdaServer_mo_to_stockin_tool"
}

type SCDDEntry struct {
	FID             int `json:",,string"`
	FBILLNO         string
	FENTRYID        int `json:",,string"`
	FSEQ            int `json:",,string"`
	FNUMBER         string
	FNAME           string
	FSPECIFICATION  string
	FBaseUnitNumber string
	FLOT_TEXT       string
	FMustQty        string
	SQTY            string
	FUseOrgNumber   string
	FPrice          string              `xorm:"-"`
	FStockNumber    string              `xorm:"-"`
	FNote           string              `xorm:"-"`
	FStockStatusId  string              `xorm:"-"`
	FSrcBillType    string              `xorm:"-"`
	FLinkInfo       []map[string]string `xorm:"-" json:"-"`
	FISBATCHMANAGE  string
}

func (*SCDDEntry) TableName() string {
	return "xkPdaServer_mo_to_stockin_tool"
}

//func GetAllSCDDMain(orgNumber string) []*SCDDMain {
//	return getSCDDMain(orgNumber, "", "")
//}

func GetSCDDMain(orgNumber, fNumber, fBillNo string) []*SCDDMain {
	return getSCDDMain(orgNumber, fNumber, fBillNo)
}

func getSCDDMain(orgNumber, fNumber, fBillNo string) (r []*SCDDMain) {
	siss := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNumber))
	fmt.Println("fBillNo:" + fBillNo)
	if fBillNo != "" {
		siss = siss.And(fmt.Sprintf("FBILLNO like '%s%%'", fBillNo))
	}
	fmt.Println("fNumber:" + fNumber)
	if fNumber != "" {
		siss = siss.And(fmt.Sprintf("FNUMBER like '%s%%'", fNumber))
	}
	e := siss.Distinct("FBILLNO, FUseOrgNumber, FNUMBER").Limit(500).Find(&r)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	parseSQL, _ := siss.LastSQL()
	fmt.Println(parseSQL)
	if len(r) == 0 {
		fmt.Println("返回nil")
		return nil
	}

	return r
}

func GetSCDDEntry(FBillNo string) []*SCDDEntry {
	return getSCDDEntry(FBillNo)
}

func getSCDDEntry(FBillNo string) (r []*SCDDEntry) {
	e := db.Where(fmt.Sprintf("FBILLNO = '%s'", FBillNo)).Limit(500).Find(&r)
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
