package dbTools

type QTCRKEntry struct {
	FUseOrgNumber   string
	FCustName       string `json:"-"`
	FCustNumber     string
	FNumber         string
	FName           string
	FSPECIFICATION  string
	FBaseUnitNumber string
	FQTY            string
	FPrice          string
	FStockNumber    string
	FStockStatusId  string `json:"-"`
	FLOT_TEXT       string
}
