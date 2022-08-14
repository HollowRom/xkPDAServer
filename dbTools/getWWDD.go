package dbTools

import (
	"fmt"
)

type WWDDMain struct {
	FBILLNO         string
	FSupplierNumber string
	FSupplierName   string
	FUseOrgNumber   string
}

func (*WWDDMain) TableName() string {
	return "xkPdaServer_pur_to_wwstockin_tool"
}

type WWDDEntry struct {
	FID             int `json:",,string"`
	FBILLNO         string
	FSupplierName   string
	FSupplierNumber string
	FENTRYID        int `json:",,string"`
	FSEQ            int `json:",,string"`
	FMATERIALID     int `json:",,string"`
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
	FStockStatusId  string              `xorm:"-"`
	FKeeperId       string              `xorm:"-"`
	FLinkInfo       []map[string]string `xorm:"-" json:"-"`
	FSrcBillNo      string              //收料通知单
	FISBATCHMANAGE string
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
		siss = siss.And(fmt.Sprintf("FSupplierNumber = '%s'", supplierNumber))
	}
	if FBillNo != "" {
		siss = siss.And(fmt.Sprintf("FBILLNO like '%s%%'", FBillNo))
	}
	e := siss.GroupBy("FBILLNO, FSupplierNumber, FSupplierName, FUseOrgNumber").Find(&r)
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
