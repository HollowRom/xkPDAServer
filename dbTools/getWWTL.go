package dbTools

import (
	"fmt"
)

type WWTLMain struct {
	FBILLNO       string
	FSuppNumber   string
	FSuppName     string
	FUseOrgNumber string
}

func (*WWTLMain) TableName() string {
	return "xkPdaServer_sub_ppbom_to_stockout_tool"
}

type WWTLEntry struct {
	FID             int `json:",,string"`
	FBILLNO         string
	FParentNumber   string
	FParentName     string
	FSuppNumber     string
	FSuppName       string
	FENTRYID        int `json:",,string"`
	FSEQ            int `json:",,string"`
	FSUBREQBILLNO   string
	FSUBREQID       int `json:",,string"`
	FSUBREQENTRYSEQ int `json:",,string"`
	FSUBREQENTRYID  int `json:",,string"`
	FMATERIALID     int `json:",,string"`
	FNUMBER         string
	FNAME           string
	FBaseUnitNumber string
	FLOT_TEXT       string
	FMustQty        string
	SQTY            string
	FUseOrgNumber   string
	FStockNumber    string              `xorm:"-"`
	FStockStatusId  string              `xorm:"-"`
	FLinkInfo       []map[string]string `xorm:"-" json:"-"`
	FSrcInterId     int
	FSrcEntryId     int
	FSrcBillNo      string
	FSrcEntrySeq    int
	FPOOrderBillNo  string
	FISBATCHMANAGE string
}

func (*WWTLEntry) TableName() string {
	return "xkPdaServer_sub_ppbom_to_stockout_tool"
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
