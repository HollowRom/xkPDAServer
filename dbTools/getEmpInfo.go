package dbTools

import (
	"fmt"
)

type EmpInfo struct {
	FMASTERID     int
	FNUMBER       string
	FNAME         string
	FUSEORGID     int
	FUseOrgNumber string
}

func (*EmpInfo) TableName() string {
	return "xkPdaServer_empInfo_tool"
}

func GetAllEmp(orgNum string) []*EmpInfo {
	return getEmp(orgNum, "")
}

func GetEmp(orgNum, number string) []*EmpInfo {
	return getEmp(orgNum, number)
}

func getEmp(orgNum, number string) (r []*EmpInfo) {
	if orgNum == "" {
		return nil
	}
	ssis := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNum))
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
