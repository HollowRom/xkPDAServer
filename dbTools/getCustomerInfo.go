package dbTools

import (
	"fmt"
)

type CustomerInfo struct {
	FMASTERID  int
	FNUMBER    string
	FNAME      string
	FUSEORGID  int
	ForgNumber string
}

func (*CustomerInfo) TableName() string {
	return "xkPdaServer_customer_tool"
}

func GetAllCustomer(orgNum string) []*CustomerInfo {
	return getCustomer(orgNum, "")
}

func GetCustomer(orgNum, number string) []*CustomerInfo {
	return getCustomer(orgNum, number)
}

func getCustomer(orgNum, number string) (r []*CustomerInfo) {
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
