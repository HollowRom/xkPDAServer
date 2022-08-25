package dbTools

import (
	"fmt"
)

type GoodsInto struct {
	FMATERIALID     int
	FNUMBER         string
	FNAME           string
	FSPECIFICATION  string
	FERPCLSID       string
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
	return getGood(orgNum, "")
}

func GetGood(orgNum, number string) []*GoodsInto {
	return getGood(orgNum, number)
}

func getGood(orgNum, number string) (r []*GoodsInto) {
	if orgNum == "" {
		return nil
	}
	siss := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNum))

	if number != "" {
		siss = siss.And(fmt.Sprintf(" (FNUMBER like '%s%%' or FNAME like '%s%%') ", number, number))
	}

	e := siss.Limit(500).Find(&r)
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
