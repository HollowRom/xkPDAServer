package dbTools

import (
	"encoding/json"
	"fmt"
)

type OrgInfo struct {
	FNumber string
	FNAME   string
	FORGID  string
}

func GetAllOrg() []*OrgInfo {
	return getOrg("")
}

func GetOrg(number string) []*OrgInfo {
	return getOrg(number)
}

func getOrg(number string) []*OrgInfo {
	selSQL := GetOrgInfo
	if number != "" {
		selSQL += " and (a.FNumber like '%" + number + "%' or b.FName like '%" + number + "%')"
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

	var rs []*OrgInfo

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
