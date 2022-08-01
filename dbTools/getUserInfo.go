package dbTools

import (
	"encoding/json"
	"fmt"
)

type UserInfo struct {
	FUSERID string
	FNAME   string
}

func GetAllUser() []*UserInfo {
	return getUser("")
}

func GetUser(number string) []*UserInfo {
	return getUser(number)
}

func getUser(number string) []*UserInfo {
	selSQL := GetUserInfo
	if number != "" {
		selSQL += " and (FName like '%" + number + "%' or FUSERID like '%" + number + "%')"
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

	var rs []*UserInfo

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
