package dbTools

import (
	"encoding/json"
	"fmt"
)

type GoodsInto struct {
	FMASTERID      string
	GoodFNUMBER    string
	FNAME          string
	FSPECIFICATION string
	FSTOCKID       string
	UnitFNUMBER    string
}

func GetAllGood(orgNum string) []*GoodsInto {
	return getGood("", orgNum)
}

func GetGood(number string, orgNum string) []*GoodsInto {
	return getGood(number, orgNum)
}

func getGood(number string, orgNum string) []*GoodsInto {
	selSQL := fmt.Sprintf(GetGoods, orgNum)
	if number != "" {
		selSQL += " and (a.FNumber like '%" + number + "%' or d.FName like '%" + number + "%')"
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

	var rs []*GoodsInto

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
