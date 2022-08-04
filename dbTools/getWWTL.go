package dbTools

import (
	"encoding/json"
	"fmt"
)

type WWTLMain struct {
	FBillNo         string
	FSupplierNumber string
	FSupplierName   string
}

type WWTLEntry struct {
	FID         int
	FENTRYID    int
	FItemNumber int
	FUnitNumber int
	FQTY        string
}

func GetAllWWTLMain(orgNumber string) []*WWTLMain {
	return getWWTLMain(orgNumber, "", "")
}

func GetWWTLMain(orgNumber, supplierNumber, FBillNo string) []*WWTLMain {
	return getWWTLMain(orgNumber, supplierNumber, FBillNo)
}

func getWWTLMain(orgNumber, supplierNumber, FBillNo string) []*WWTLMain {
	if orgNumber == "" {
		return nil
	}
	sql := fmt.Sprintf(GetWWTLMainInfo, orgNumber)
	if supplierNumber != "" {
		sql += " and (f.FNUMBER = '" + supplierNumber + "' or g.FNAME = '" + supplierNumber + "')"
	}
	if FBillNo != "" {
		sql += " and a.FBILLNO like '%" + FBillNo + "%'"
	}
	r, e := db.QueryString(sql)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	if len(r) == 0 {
		fmt.Println("返回nil")
		return nil
	}

	j, e := json.Marshal(r)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	var rs []*WWTLMain

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}

func GetWWTLEntry(FBillNo string) []*WWTLEntry {
	return getWWTLEntry(FBillNo)
}

func getWWTLEntry(FBillNo string) []*WWTLEntry {
	if FBillNo == "" {
		return nil
	}
	r, e := db.QueryString(fmt.Sprintf(GetWWTLEntryInfo, FBillNo))
	if e != nil {
		fmt.Println(e)
		return nil
	}

	if len(r) == 0 {
		fmt.Println("返回nil")
		return nil
	}

	j, e := json.Marshal(r)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	var rs []*WWTLEntry

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
