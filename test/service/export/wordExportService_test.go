package export

import (
	"fmt"
	"github.com/Casper-Mars/dbTool/service/db"
	"github.com/Casper-Mars/dbTool/service/export"
	"testing"
)

var username_mssql = "sa"
var password_mssql = "!Zhisheng2020"
var ipPort_mssql = "192.168.123.156:1434"
var dbName_mssql = "DC_ZJDDGL"

func TestExport(t *testing.T) {

	dbService := db.MsService{}
	info, err := dbService.GetAllTableInfo(username_mssql, password_mssql, ipPort_mssql, dbName_mssql)
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Printf("%#v", info)
	}
	service := export.WordExportService{}
	service.Export(dbName_mssql, "./", info)

}
