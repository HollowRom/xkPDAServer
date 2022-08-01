package dbTools

import (
	"encoding/json"
	"fmt"
)

type CGDDMain struct {
	FBillNo         string
	FSupplierNumber string
	FSupplierName   string
}

type CGDDEntry struct {
	FID         int
	FENTRYID    int
	FItemNumber int
	FUnitNumber int
	FQTY        string
}

func GetAllCGDDMain(orgNumber string) []*CGDDMain {
	return getCGDDMain(orgNumber, "", "")
}

func GetCGDDMain(orgNumber, supplierNumber, FBillNo string) []*CGDDMain {
	return getCGDDMain(orgNumber, supplierNumber, FBillNo)
}

func getCGDDMain(orgNumber, supplierNumber, FBillNo string) []*CGDDMain {
	if orgNumber == "" {
		return nil
	}
	sql := fmt.Sprintf(GetCGDDMainInfo, orgNumber)
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

	var rs []*CGDDMain

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}

func GetCGDDEntry(FBillNo string) []*CGDDEntry {
	return getCGDDEntry(FBillNo)
}

func getCGDDEntry(FBillNo string) []*CGDDEntry {
	if FBillNo == "" {
		return nil
	}
	r, e := db.QueryString(fmt.Sprintf(GetCGDDEntryInfo, FBillNo))
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

	var rs []*CGDDEntry

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
