package dbTools

import (
	"fmt"
)

type CGDDMain struct {
	FBILLNO       string
	FSuppNumber   string
	FSuppName     string
	FUseOrgNumber string
}

func (*CGDDMain) TableName() string {
	return "xkPdaServer_sltz_to_cgrk_tool"
}

type CGDDEntry struct {
	FID             int `json:",,string"`
	FBILLNO         string
	FSUPPLIERID     int `json:",,string"`
	FSuppNumber     string
	FSuppName       string
	FPURCHASERID    int `json:",,string"`
	FSRCBILLNO      string
	FSRCENTRYID     int `json:",,string"`
	FSRCID          int  `json:",,string"`
	FSRCSEQ         int `json:",,string"`
	FMATERIALID     int `json:",,string"`
	FENTRYID        int `json:",,string"`
	FSEQ            int  `json:",,string"`
	FNUMBER         string
	FNAME           string
	FSPECIFICATION  string
	FBaseUnitNumber string
	FLOT_TEXT       string
	FMustQty        string
	SQTY string
	FUseOrgNumber   string
	FSrcBillType    string              `xorm:"-"`
	FKeeperId       string              `xorm:"-"`
	FPrice          string              `xorm:"-"`
	FStockNumber    string              `xorm:"-"`
	FNote           string              `xorm:"-"`
	FStockStatusId  string              `xorm:"-"`
	FLinkInfo       []map[string]string `xorm:"-" json:"-"`
	FISBATCHMANAGE string
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
