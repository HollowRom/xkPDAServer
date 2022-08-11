package dbTools

import (
	"fmt"
)

type WWDDMain struct {
	FBillNo     string
	FSuppNumber string
	FSuppName   string
}

func (*WWDDMain) TableName() string {
	return "xkPdaServer_pur_to_wwstockin_tool"
}

type WWDDEntry struct {
	FID             int
	FBILLNO         string
	FSupplierName   string
	FSupplierNumber string
	FENTRYID        int
	FSEQ            int
	FMATERIALID     int
	FNUMBER         string
	FNAME           string
	FSPECIFICATION  string
	FBaseUnitNumber string
	FLOT_TEXT       string
	FMustQty        string
	SQTY            string
	FUseOrgNumber   string
}

func (*WWDDEntry) TableName() string {
	return "xkPdaServer_pur_to_wwstockin_tool"
}

func GetAllWWDDMain(orgNumber string) []*WWDDMain {
	return getWWDDMain(orgNumber, "", "")
}

func GetWWDDMain(orgNumber, supplierNumber, FBillNo string) []*WWDDMain {
	return getWWDDMain(orgNumber, supplierNumber, FBillNo)
}

func getWWDDMain(orgNumber, supplierNumber, FBillNo string) (r []*WWDDMain) {
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

func GetWWDDEntry(FBillNo string) []*WWDDEntry {
	return getWWDDEntry(FBillNo)
}

func getWWDDEntry(FBillNo string) (r []*WWDDEntry) {
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
