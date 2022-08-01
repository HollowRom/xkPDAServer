package dbTools

import (
	"encoding/json"
	"fmt"
)

type KeeperInfo struct {
	FNumber string
	FNAME   string
	FItemId string
}

func GetAllKeeper(orgNum string) []*KeeperInfo {
	return getKeeper("", orgNum)
}

func GetKeeper(number string, orgNum string) []*KeeperInfo {
	return getKeeper(number, orgNum)
}

func getKeeper(number string, orgNum string) []*KeeperInfo {
	selSQL := fmt.Sprintf(GetKeeperInfo, orgNum)
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

	var rs []*KeeperInfo

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
