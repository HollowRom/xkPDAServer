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

func GetAllCustomer() []*CustomerInfo {
	return getCustomer("")
}

func GetCustomer(number string) []*CustomerInfo {
	return getCustomer(number)
}

func getCustomer(number string) (r []*CustomerInfo) {
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
