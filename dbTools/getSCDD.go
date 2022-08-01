package dbTools

import (
	"encoding/json"
	"fmt"
)

type SCDDMain struct {
	FBillNo string
}

type SCDDEntry struct {
	FNumber      string
	FSrcEntryId  int
	UnitNumber   string
	FQTY         string
	FPrice       string
	FStockNumber string
	FNote        string
	FLot         string
	FMoBillNo    string
	FMoId        int
	FMoEntryId   string
	FMoEntrySeq  int
	FSrcBillNo   string
	FSrcBillType string
	FSrcInterId  int
}

func GetAllSCDDMain(orgNumber string) []*SCDDMain {
	return getSCDDMain(orgNumber, "")
}

func GetSCDDMain(orgNumber, fBillNo string) []*SCDDMain {
	return getSCDDMain(orgNumber, fBillNo)
}

func getSCDDMain(orgNumber, fBillNo string) []*SCDDMain {
	if orgNumber == "" {
		return nil
	}
	sql := fmt.Sprintf(GetSCDDMainInfo, orgNumber)
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

	var rs []*SCDDMain

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}

func GetSCDDEntry(FBillNo string) []*SCDDEntry {
	return getSCDDEntry(FBillNo)
}

func getSCDDEntry(FBillNo string) []*SCDDEntry {
	if FBillNo == "" {
		return nil
	}
	sql := fmt.Sprintf(GetSCDDEntryInfo, FBillNo)
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

	var rs []*SCDDEntry

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
