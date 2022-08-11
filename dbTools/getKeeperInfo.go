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

func GetAllKeeper(orgNum string) []*KeeperInfo {
	return getKeeper("", orgNum)
}

func GetKeeper(number string, orgNum string) []*KeeperInfo {
	return getKeeper(number, orgNum)
}

func getKeeper(number string, orgNum string) (r []*KeeperInfo) {
	e := db.Where(fmt.Sprintf("(FName = '%s' or FNumber = '%s') and FUseOrgNumber = '%s' ", number, number, orgNum)).Find(&r)
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
