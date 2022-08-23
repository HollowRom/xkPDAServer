package dbTools

type QTCRKEntry struct {
	FUseOrgNumber   string `xorm:"-" json:"FKeeperId"`
	FCustName       string `json:"-"` //合并了出入库单所以供应商和客户分不开
	FCustNumber     string //合并了出入库单所以供应商和客户分不开
	FNumber         string
	FName           string
	FSPECIFICATION  string `json:"-"`
	FBaseUnitNumber string
	FQTY            string
	FPrice          string
	FStockNumber    string
	FStockStatusId  string `json:"-"`
	FLOT_TEXT       string
}
