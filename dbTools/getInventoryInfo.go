package dbTools

import "fmt"

type InventoryInfo struct {
	FUseOrgNumber string
	FStockNumber string
	FStockName string
	FLot_Text string
	FMATERIALID int `json:",,string"`
	FNUMBER string
	FNAME string
	FBASEQTY float64 `json:",,string"`
	FBaseUnitNumber string
	FUPDATETIME string
}


func (*InventoryInfo) TableName() string {
	return "xkPdaServer_inventory_tool"
}

//func GetAllInventory(orgNum string) []*InventoryInfo {
//	return getInventory(orgNum, "", "", "")
//}

func GetInventory(orgNum, fNumber, lotNumber, stockNumber string) []*InventoryInfo {
	return getInventory(orgNum, fNumber, lotNumber, stockNumber)
}

func getInventory(orgNum, fNumber, lotNumber, stockNumber string) (r []*InventoryInfo) {
	if orgNum == "" {
		return nil
	}
	ssis := db.Where(fmt.Sprintf("FUseOrgNumber = '%s'", orgNum))
	if fNumber != "" {
		ssis = ssis.And(fmt.Sprintf(" FNUMBER like '%s%%' ", fNumber))
	}
	if lotNumber != "" {
		ssis = ssis.And(fmt.Sprintf(" FLot_Text like '%s%%' ", lotNumber))
	}
	if stockNumber != "" {
		ssis = ssis.And(fmt.Sprintf(" FStockNumber like '%s%%' ", stockNumber))
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

