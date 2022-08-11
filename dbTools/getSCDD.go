package dbTools

import (
	"fmt"
)

type SCDDMain struct {
	FBILLNO string
	FNumber string
	FName   string
}

func (*SCDDMain) TableName() string {
	return "xkPdaServer_mo_to_stockin_tool"
}

type SCDDEntry struct {
	FID             int
	FBILLNO         string
	FENTRYID        int
	FSEQ            int
	FNUMBER         string
	FNAME           string
	FSPECIFICATION  string
	FBaseUnitNumber string
	FLOT_TEXT       string
	FMustQty        string
	SQTY            string
	FUseOrgNumber   string
}

func (*SCDDEntry) TableName() string {
	return "xkPdaServer_mo_to_stockin_tool"
}

func GetAllSCDDMain(orgNumber string) []*SCDDMain {
	return getSCDDMain(orgNumber, "")
}

func GetSCDDMain(orgNumber, fBillNo string) []*SCDDMain {
	return getSCDDMain(orgNumber, fBillNo)
}

func getSCDDMain(orgNumber, fBillNo string) (r []*SCDDMain) {
	siss := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNumber))
	if fBillNo != "" {
		siss = siss.And(fmt.Sprintf("FBILLNO like '%s%%'", fBillNo))
	}
	e := siss.GroupBy("FBILLNO, FNumber, FName").Find(&r)
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

func GetSCDDEntry(FBillNo string) []*SCDDEntry {
	return getSCDDEntry(FBillNo)
}

func getSCDDEntry(FBillNo string) (r []*SCDDEntry) {
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
