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
	return getEmp("", orgNum)
}

func GetEmp(number string, orgNum string) []*EmpInfo {
	return getEmp(number, orgNum)
}

func getEmp(number string, orgNum string) (r []*EmpInfo) {
	e := db.Where(fmt.Sprintf("(FName like '%%%s%%' or FNUMBER like '%%%s%%') and FUseOrgNumber = '%s'", number, number, orgNum)).Find(&r)
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
