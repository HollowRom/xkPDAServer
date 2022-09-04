package dbTools

import "fmt"

type DepartInfo struct {
	FMASTERID int `json:",,string"`
	FNUMBER string
	FNAME string
	FUseOrgNumber string
	FUSEORGID int `json:",,string"`
}

func (*DepartInfo) TableName() string {
	return "xkPdaServer_depart_tool"
}

func GetDepart(orgNum, number string) []*DepartInfo {
	return getDepart(orgNum, number)
}

func getDepart(orgNum, number string) (r []*DepartInfo) {
	if orgNum == "" {
		return nil
	}
	ssis := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNum))
	if number != "" {
		ssis = ssis.And(fmt.Sprintf(" FNUMBER like '%s%%' ", number))
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

