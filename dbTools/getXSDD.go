package dbTools

import (
	"encoding/json"
	"fmt"
)

type XSDDMain struct {
	FBillNo         string
	FCustomerNumber string
	FCustomerName   string
}

type XSDDEntry struct {
	FCustMatID   string
	FNumber      string
	UnitNumber   string
	FQTY         string
	FStockNumber string
	FPrice       string
	FSoorDerno   string
	FSOEntryId   int
	FSOInterId   int
}

func GetAllXSDDMain(orgNumber string) []*XSDDMain {
	return getXSDDMain(orgNumber, "")
}

func GetXSDDMain(orgNumber, fBillNo string) []*XSDDMain {
	return getXSDDMain(orgNumber, fBillNo)
}

func getXSDDMain(orgNumber, fBillNo string) []*XSDDMain {
	if orgNumber == "" {
		return nil
	}
	sql := fmt.Sprintf(GetXSDDMainInfo, orgNumber)
	if fBillNo != "" {
		sql += " and a.FBillNo like '%" + orgNumber + "%'"
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

	var rs []*XSDDMain

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}

func GetXSDDEntry(FBillNo string) []*XSDDEntry {
	return getXSDDEntry(FBillNo)
}

func getXSDDEntry(FBillNo string) []*XSDDEntry {
	if FBillNo == "" {
		return nil
	}
	sql := fmt.Sprintf(GetXSDDEntryInfo, FBillNo)
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

	var rs []*XSDDEntry

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
