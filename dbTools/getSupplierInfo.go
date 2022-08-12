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

func GetAllSupplier() []*SupplierInfo {
	return getSupplier("")
}

func GetSupplier(number string) []*SupplierInfo {
	return getSupplier(number)
}

func getSupplier(number string) (r []*SupplierInfo) {
	e := db.Where(fmt.Sprintf("FName = '%s' or FNUMBER = '%s'", number, number)).Find(&r)
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
