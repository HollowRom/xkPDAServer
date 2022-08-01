package dbTools

import (
	"encoding/json"
	"fmt"
)

type SupplierInfo struct {
	FMASTERID string
	FNUMBER   string
	FNAME     string
}

func GetAllSupplier() []*SupplierInfo {
	return getSupplier("")
}

func GetSupplier(number string) []*SupplierInfo {
	return getSupplier(number)
}

func getSupplier(number string) []*SupplierInfo {
	selSQL := GetSupplierInfo
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

	var rs []*SupplierInfo

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
