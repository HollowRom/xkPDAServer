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

func GetAllStock(orgNum string) []*StockInto {
	return getStock("", orgNum)
}

func GetStock(number string, orgNum string) []*StockInto {
	return getStock(number, orgNum)
}

func getStock(number string, orgNum string) (r []*StockInto) {
	e := db.Where(fmt.Sprintf("FName = '%s' or FNUMBER = '%s'", number, number)).Find(&r)
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
