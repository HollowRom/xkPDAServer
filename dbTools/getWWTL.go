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
	FID             int
	FBILLNO         string
	FParentNumber   string
	FParentName     string
	FSuppNumber     string
	FSuppName       string
	FENTRYID        int
	FSEQ            int
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
	FStockNumber    string              `xorm:"-"`
	FStockStatusId  string              `xorm:"-"`
	FKeeperId       int                 `xorm:"-"`
	FLinkInfo       []map[string]string `xorm:"-"`
	FSrcInterId     int
	FSrcEntryId     int
	FSrcBillNo      string
	FSrcEntrySeq    int
	FPOOrderBillNo  string
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
