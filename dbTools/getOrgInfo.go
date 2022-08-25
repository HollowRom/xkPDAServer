package dbTools

import (
	"fmt"
)

type OrgInfo struct {
	FNumber string
	FNAME   string
	FORGID  string
}

func (*OrgInfo) TableName() string {
	return "xkPdaServer_orgInfo_tool"
}

func GetAllOrg() []*OrgInfo {
	return getOrg("")
}

func GetOrg(number string) []*OrgInfo {
	return getOrg(number)
}

func getOrg(number string) (r []*OrgInfo) {
	e := db.Where(fmt.Sprintf("FName like '%%%s%%' or FNUMBER like '%%%s%%'", number, number)).Limit(500).Find(&r)
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
