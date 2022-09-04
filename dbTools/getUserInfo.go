package dbTools

import (
	"fmt"
)

type UserInfo struct {
	FUSERID string
	FNAME   string
}

func (*UserInfo) TableName() string {
	return "xkPdaServer_userInfo_tool"
}

//func GetAllUser() []*UserInfo {
//	return getUser("")
//}

func GetUser(number string) []*UserInfo {
	return getUser(number)
}

func getUser(number string) (r []*UserInfo) {
	e := db.Where(fmt.Sprintf("FName = '%s' or FUSERID = %s", number, number)).Limit(500).Find(&r)
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
