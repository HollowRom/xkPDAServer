package dbTools

import (
	"fmt"
)

type KeeperInfo struct {
	FNumber       string
	FNAME         string
	FKeeperId     string
	FORGID        int
	FUseOrgNumber string
}

func (*KeeperInfo) TableName() string {
	return "xkPdaServer_keeperInfo_tool"
}

//func GetAllKeeper(orgNum string) []*KeeperInfo {
//	return getKeeper(orgNum, "")
//}

func GetKeeper(orgNum, number string) []*KeeperInfo {
	return getKeeper(orgNum, number)
}

func getKeeper(orgNum, number string) (r []*KeeperInfo) {
	e := db.Where(fmt.Sprintf("(FName = '%s' or FNUMBER = '%s') and FUseOrgNumber = '%s' ", number, number, orgNum)).Limit(500).Find(&r)
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
