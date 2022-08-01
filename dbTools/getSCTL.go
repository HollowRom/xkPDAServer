package dbTools

import (
	"encoding/json"
	"fmt"
)

type SCTLMain struct {
	FBillNo       string
	FParentNumber string
}

type SCTLEntry struct {
	FParentNumber     string
	FNumber           string
	UnitNumber        string
	FQTY              string
	FStockNumber      string
	FMoBillNo         string
	FMoEntryId        int
	FPPBomEntryId     int
	FMoId             int
	FMoEntrySeq       int
	FPPBomBillNo      string
	FEntrySrcInterId  int
	FEntrySrcEntrySeq int
	FKeeperId         string
	FLotNo            string
}

func GetAllSCTLMain(orgNumber string) []*SCTLMain {
	return getSCTLMain(orgNumber, "")
}

func GetSCTLMain(orgNumber, fBillNo string) []*SCTLMain {
	return getSCTLMain(orgNumber, fBillNo)
}

func getSCTLMain(orgNumber, fBillNo string) []*SCTLMain {
	if orgNumber == "" {
		return nil
	}
	sql := fmt.Sprintf(GetSCTLMainInfo, orgNumber)
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

	var rs []*SCTLMain

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}

func GetSCTLEntry(FBillNo string) []*SCTLEntry {
	return getSCTLEntry(FBillNo)
}

func getSCTLEntry(FBillNo string) []*SCTLEntry {
	if FBillNo == "" {
		return nil
	}
	sql := fmt.Sprintf(GetSCTLEntryInfo, FBillNo)
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

	var rs []*SCTLEntry

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
