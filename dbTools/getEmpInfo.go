package dbTools

import (
	"encoding/json"
	"fmt"
)

type EmpInfo struct {
	FMASTERID string
	FNUMBER   string
	FNAME     string
}

func GetAllEmp(orgNum string) []*EmpInfo {
	return getEmp("", orgNum)
}

func GetEmp(number string, orgNum string) []*EmpInfo {
	return getEmp(number, orgNum)
}

func getEmp(number string, orgNum string) []*EmpInfo {
	selSQL := fmt.Sprintf(GetEmpInfo, orgNum)
	if number != "" {
		selSQL += " and (a.FNumber like '%" + number + "%' or b.FName like '%" + number + "%')"
	}
	r, e := db.QueryString(selSQL)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	if len(r) == 0 {
		return nil
	}

	j, e := json.Marshal(r)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	var rs []*EmpInfo

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
