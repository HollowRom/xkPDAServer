package dbTools

import (
	"encoding/json"
	"fmt"
)

type WWDDMain struct {
	FBillNo         string
	FSupplierNumber string
	FSupplierName   string
}

type WWDDEntry struct {
	FID         int
	FENTRYID    int
	FItemNumber int
	FUnitNumber int
	FQTY        string
}

func GetAllWWDDMain(orgNumber string) []*WWDDMain {
	return getWWDDMain(orgNumber, "", "")
}

func GetWWDDMain(orgNumber, supplierNumber, FBillNo string) []*WWDDMain {
	return getWWDDMain(orgNumber, supplierNumber, FBillNo)
}

func getWWDDMain(orgNumber, supplierNumber, FBillNo string) []*WWDDMain {
	if orgNumber == "" {
		return nil
	}
	sql := fmt.Sprintf(GetWWDDMainInfo, orgNumber)
	if supplierNumber != "" {
		sql += " and (f.FNUMBER = " + supplierNumber + " or g.FNAME = " + supplierNumber + ")"
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

	var rs []*WWDDMain

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}

func GetWWDDEntry(FBillNo string) []*WWDDEntry {
	return getWWDDEntry(FBillNo)
}

func getWWDDEntry(FBillNo string) []*WWDDEntry {
	if FBillNo == "" {
		return nil
	}
	r, e := db.QueryString(fmt.Sprintf(GetWWDDEntryInfo, FBillNo))
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

	var rs []*WWDDEntry

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
