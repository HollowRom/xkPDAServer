package dbTools

import (
	"fmt"
)

type GoodsInto struct {
	FMATERIALID     int
	FNUMBER         string
	FNAME           string
	FSPECIFICATION  string
	FERPCLSID       int
	FErpClass       string
	FSUITE          byte
	FBASEUNITID     int
	FBaseUnitNumber string
	FWEIGHTUNITID   int
	FVOLUMEUNITID   int
	FISPURCHASE     byte
	FISINVENTORY    byte
	FISSUBCONTRACT  byte
	FISSALE         byte
	FISPRODUCE      byte
	FISASSET        byte
	FISBATCHMANAGE  byte
	FISKFPERIOD     byte
	FCHECKINCOMING  byte
	FCHECKPRODUCT   byte
	FCHECKSTOCK     byte
	FCHECKRETURN    byte
	FCHECKDELIVERY  byte
	FSTOCKID        int
	FStockName      string
	FUSEORGID       int
	FUseOrgNumber   string
	FUseOrgName     string
}

func (*GoodsInto) TableName() string {
	return "xkPdaServer_good_tool"
}

func GetAllGood(orgNum string) []*GoodsInto {
	return getGood("", orgNum)
}

func GetGood(number string, orgNum string) []*GoodsInto {
	return getGood(number, orgNum)
}

func getGood(number string, orgNum string) (r []*GoodsInto) {
	siss := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNum))

	if number != "" {
		siss = siss.And(fmt.Sprintf(" (FNUMBER like '%s%%' or FNAME like '%s%%') ", number, number))
	}

	e := siss.Find(&r)
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
