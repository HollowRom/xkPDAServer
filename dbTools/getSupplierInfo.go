package dbTools

import (
	"fmt"
)

type SupplierInfo struct {
	FMASTERID        int
	FNUMBER          string
	FNAME            string
	FSUPPLYCLASSIFY  string
	FSUPPLYCLASSName string
	FORGID           int
	ForgNumber       string
}

func (*SupplierInfo) TableName() string {
	return "xkPdaServer_supplier_tool"
}

//func GetAllSupplier(orgNum string) []*SupplierInfo {
//	return getSupplier(orgNum, "")
//}

func GetSupplier(orgNum, number string) []*SupplierInfo {
	return getSupplier(orgNum, number)
}

func getSupplier(orgNum, number string) (r []*SupplierInfo) {
	if orgNum == "" {
		return nil
	}
	ssis := db.Where(fmt.Sprintf("ForgNumber = '%s'", orgNum))
	if number != "" {
		ssis = ssis.And(fmt.Sprintf(" (FName like '%s%%' or FNUMBER like '%s%%') ", number, number))
	}
	e := ssis.Limit(500).Find(&r)
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
