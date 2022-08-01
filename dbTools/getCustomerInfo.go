package dbTools

import (
	"encoding/json"
	"fmt"
)

type CustomerInfo struct {
	FMASTERID string
	FNUMBER   string
	FNAME     string
}

func GetAllCustomer() []*CustomerInfo {
	return getCustomer("")
}

func GetCustomer(number string) []*CustomerInfo {
	return getCustomer(number)
}

func getCustomer(number string) []*CustomerInfo {
	selSQL := GetCustomerInfo
	if number != "" {
		selSQL += " and (b.FName like '%" + number + "%' or a.FNumber like '%" + number + "%')"
	}
	r, e := db.QueryString(selSQL)
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

	var rs []*CustomerInfo

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
