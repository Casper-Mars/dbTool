package db

import (
	"fmt"
	"github.com/Casper-Mars/dbTool/service/db"
	"testing"
)

var username_mssql = "sa"
var password_mssql = "!Zhisheng2020"
var ipPort_mssql = "192.168.123.156:1434"
var dbName_mssql = "DC_ZJDDGL"

var username_mysql = "root"
var password_mysql = "!Zhisheng2020"
var ipPort_mysql = "192.168.123.155:3306"
var dbName_mysql = "taxi_authority"

func TestGetAllTableInfoMssql(t *testing.T) {
	dbService := db.MsService{}
	info, err := dbService.GetAllTableInfo(username_mssql, password_mssql, ipPort_mssql, dbName_mssql)
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Printf("%#v", info)
	}
}
func TestGetAllTableInfoMysql(t *testing.T) {
	dbService := db.MysqlService{}
	info, err := dbService.GetAllTableInfo(username_mysql, password_mysql, ipPort_mysql, dbName_mysql)
	if err != nil {
		fmt.Println(err)
		panic(err)
	} else {
		fmt.Printf("%#v", info)
	}
}

func TestGetAllDBsMssql(t *testing.T) {
	dbService := db.MsService{}
	bs, err := dbService.GetAllDBs(username_mssql, password_mssql, ipPort_mssql)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%#v", bs)
	}
}
func TestGetAllDBsMysql(t *testing.T) {
	dbService := db.MysqlService{}
	bs, err := dbService.GetAllDBs(username_mysql, password_mysql, ipPort_mysql)
	if err != nil {
		panic(err)
	} else {
		fmt.Printf("%#v", bs)
	}
}
