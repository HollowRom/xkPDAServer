package dbTools

import (
	"fmt"
)

type CGDDMain struct {
	FBILLNO     string
	FSuppNumber string
	FSuppName   string
}

func (*CGDDMain) TableName() string {
	return "xkPdaServer_sltz_to_cgrk_tool"
}

type CGDDEntry struct {
	FID             int
	FBILLNO         string
	FSUPPLIERID     int
	FSuppNumber     string
	FSuppName       string
	FPURCHASERID    int
	FSRCBILLNO      string
	FSRCENTRYID     int
	FSRCID          int
	FSRCSEQ         int
	FMATERIALID     int
	FENTRYID        int
	FSEQ            int
	FNAME           string
	FSPECIFICATION  string
	FBaseUnitNumber int
	FLOT_TEXT       string
	FMustQty        string
	SQTY            string
}

func (*CGDDEntry) TableName() string {
	return "xkPdaServer_sltz_to_cgrk_tool"
}

func GetAllCGDDMain(orgNumber string) []*CGDDMain {
	return getCGDDMain(orgNumber, "", "")
}

func GetCGDDMain(orgNumber, supplierNumber, FBillNo string) []*CGDDMain {
	return getCGDDMain(orgNumber, supplierNumber, FBillNo)
}

func getCGDDMain(orgNumber, supplierNumber, FBillNo string) (r []*CGDDMain) {
	siss := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNumber))
	if supplierNumber != "" {
		siss = siss.And(fmt.Sprintf("FSuppNumber = '%s'", supplierNumber))
	}
	if FBillNo != "" {
		siss = siss.And(fmt.Sprintf("FBILLNO like '%s%%'", FBillNo))
	}
	e := siss.GroupBy("FBILLNO, FSuppNumber, FSuppName").Find(&r)
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

func GetCGDDEntry(FBillNo string) []*CGDDEntry {
	return getCGDDEntry(FBillNo)
}

func getCGDDEntry(FBillNo string) (r []*CGDDEntry) {
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
