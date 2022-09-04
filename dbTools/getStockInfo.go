package dbTools

import (
	"fmt"
)

type StockInto struct {
	FMASTERID  int
	FNUMBER    string
	FNAME      string
	FUSEORGID  int
	ForgNumber string
}

func (*StockInto) TableName() string {
	return "xkPdaServer_stock_tool"
}

//func GetAllStock(orgNum string) []*StockInto {
//	return getStock(orgNum, "")
//}

func GetStock(orgNum, number string) []*StockInto {
	return getStock(orgNum, number)
}

func getStock(orgNum, number string) (r []*StockInto) {
	if orgNum == "" {
		return nil
	}
	ssis := db.Where(fmt.Sprintf("ForgNumber = '%s'", orgNum))
	if number != "" {
		ssis = ssis.And(fmt.Sprintf(" (FNAME like '%s%%' or FNUMBER like '%s%%') ", number, number))
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
