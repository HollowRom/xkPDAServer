package dbTools

import (
	"fmt"
)

type WWTLMain struct {
	FBillNo     string
	FSuppNumber string
	FSuppName   string
}

func (*WWTLMain) TableName() string {
	return "xkPdaServer_sub_ppbom_to_stockin_tool"
}

type WWTLEntry struct {
	FID             int
	FBILLNO         string
	FSuppNumber     string
	FSuppName       string
	FENTRYID        int
	FSUBREQBILLNO   string
	FSUBREQID       int
	FSUBREQENTRYSEQ int
	FSUBREQENTRYID  int
	FMATERIALID     int
	FNUMBER         string
	FNAME           string
	FBaseUnitNumber string
	FLOT_TEXT       string
	FMustQty        string
	SQTY            string
	FUseOrgNumber   string
}

func (*WWTLEntry) TableName() string {
	return "xkPdaServer_sub_ppbom_to_stockin_tool"
}

func GetAllWWTLMain(orgNumber string) []*WWTLMain {
	return getWWTLMain(orgNumber, "", "")
}

func GetWWTLMain(orgNumber, supplierNumber, FBillNo string) []*WWTLMain {
	return getWWTLMain(orgNumber, supplierNumber, FBillNo)
}

func getWWTLMain(orgNumber, supplierNumber, FBillNo string) (r []*WWTLMain) {
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

func GetWWTLEntry(FBillNo string) []*WWTLEntry {
	return getWWTLEntry(FBillNo)
}

func getWWTLEntry(FBillNo string) (r []*WWTLEntry) {
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
