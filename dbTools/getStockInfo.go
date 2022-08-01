package dbTools

import (
	"encoding/json"
	"fmt"
)

type StockInto struct {
	FMASTERID string
	FNUMBER   string
	FNAME     string
}

func GetAllStock(orgNum string) []*StockInto {
	return getStock("", orgNum)
}

func GetStock(number string, orgNum string) []*StockInto {
	return getStock(number, orgNum)
}

func getStock(number string, orgNum string) []*StockInto {
	selSQL := fmt.Sprintf(GetStockInfo, orgNum)
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

	var rs []*StockInto

	e = json.Unmarshal(j, &rs)
	if e != nil {
		fmt.Println(e)
		return nil
	}

	return rs
}
