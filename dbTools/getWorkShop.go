package dbTools

import "fmt"

type WorkShopMain struct {
	FID     int `json:",,string"`
	FBILLNO string
	FNUMBER string
}

func (*WorkShopMain) TableName() string {
	return "xkPdaServer_gxplan_to_gxreport_tool"
}

type WorkShopEntry struct {
	FID             int `json:",,string"`
	FBILLNO         string
	FMOID           int `json:",,string"`
	FMOENTRYSEQ     int `json:",,string"`
	FMOENTRYID      int `json:",,string"`
	FENTRYID        int `json:",,string"`
	FSEQ            int `json:",,string"`
	FDepartNumber   string
	FSEQNUMBER      string
	FUnitNumber     string
	FWCNumber       string
	FMATERIALID     int `json:",,string"`
	FNUMBER         string
	FNAME           string
	FSPECIFICATION  string
	FBaseUnitNumber string
	FLOT_TEXT       string
	FMustQty        float64 `json:",,string"`
	SQTY            float64 `json:",,string"`
	FUseOrgNumber   string
	FISBATCHMANAGE  int `json:",,string"`
}

func (*WorkShopEntry) TableName() string {
	return "xkPdaServer_gxplan_to_gxreport_tool"
}

func GetWorkShopMain(orgNumber, fBillNo, fNumber string) []*WorkShopMain {
	return getWorkShopMain(orgNumber, fBillNo, fNumber)
}

func getWorkShopMain(orgNumber, fBillNo, fNumber string) (r []*WorkShopMain) {
	siss := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNumber))
	if fBillNo != "" {
		siss = siss.And(fmt.Sprintf("FBILLNO = '%s'", fBillNo))
	}
	if fNumber != "" {
		siss = siss.And(fmt.Sprintf("FNUMBER like '%s%%'", fNumber))
	}
	e := siss.Distinct("FID, FBILLNO, FNUMBER").Limit(500).Find(&r)
	if e != nil {
		fmt.Println(e)
		return nil
	}
	fmt.Println(siss.LastSQL())
	if len(r) == 0 {
		fmt.Println("返回nil")
		return nil
	}

	return r
}

func GetWorkShopEntry(FBillNo string) []*WorkShopEntry {
	return getWorkShopEntry(FBillNo)
}

func getWorkShopEntry(FBillNo string) (r []*WorkShopEntry) {
	e := db.Where(fmt.Sprintf("FBILLNO = '%s'", FBillNo)).Limit(500).Find(&r)
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
