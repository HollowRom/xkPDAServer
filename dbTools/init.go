package dbTools

import (
	"fmt"
	"strings"
	"sync"
	"xorm.io/xorm"
	"xorm.io/xorm/names"
)
import _ "github.com/denisenkom/go-mssqldb"

var db *xorm.Engine
var maxConnNum = 5
var defDSN = "driver={SQL Server};Server=127.0.0.1;Database=AIS20210805182552;user id=sa;password=sa;"
var defDSNConf = ""

var dsn string

var o sync.Once

var oneInit = func () {
	dbConfMap := map[string]string{}
	tempValue := GetConfFromKey("ServerIp")
	if tempValue != "" {
		dbConfMap["ServerIp"] = tempValue
	}
	tempValue = GetConfFromKey("DatabaseName")
	if tempValue != "" {
		dbConfMap["DatabaseName"] = tempValue
	}
	tempValue = GetConfFromKey("DBuid")
	if tempValue != "" {
		dbConfMap["DBuid"] = tempValue
	}
	tempValue = GetConfFromKey("DBpwd")
	if tempValue != "" {
		dbConfMap["DBpwd"] = tempValue
	}

	if len(dbConfMap) == 4 {
		defDSNConf = fmt.Sprintf("driver={SQL Server};Server=%s;Database=%s;user id=%s;password=%s;", dbConfMap["ServerIp"], dbConfMap["DatabaseName"], dbConfMap["DBuid"], dbConfMap["DBpwd"])
	}

	if db != nil {
		panic("不能再次初始化db")
	}

	if defDSNConf != "" {
		dsn = defDSNConf
	}

	//奇怪的证书bug
	if !strings.Contains(dsn, "encrypt=") {
		dsn = strings.TrimRight(dsn, ";") + ";encrypt=disable"
	}

	if dsn == "" {
		dsn = defDSN
	}

	db, _ = xorm.NewEngine("mssql", dsn)

	err := db.Ping()
	if err != nil {
		panic("sqlserver ping 失败" + err.Error())
	}

	db.DB().SetMaxOpenConns(maxConnNum)

	db.SetMapper(names.SameMapper{})

	fmt.Println("db初始化完成，未出现异常")
}

func Init(DSN string) {
	dsn = DSN
	o.Do(oneInit)
}

func GetDB() *xorm.Engine {
	return db
}
