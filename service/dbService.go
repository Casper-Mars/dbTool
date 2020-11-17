package service

import (
	"github.com/Casper-Mars/dbTool/pojo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/schemas"
)

func GetAllTableInfo(username string, password string, ipPort string, dbName string) []pojo.TableInfo {
	engine := openConnect(username, password, ipPort, dbName)
	defer engine.Close()
	metas, err := engine.DBMetas()
	if err != nil {
		panic(err)
	}
	info := make([]pojo.TableInfo, len(metas))
	for i, k := range metas {
		var tableInfoTmp pojo.TableInfo
		tableInfoTmp.Comment = k.Comment
		tableInfoTmp.TableName = k.Name
		tableInfoTmp.Cols = extractCol(k)
		info[i] = tableInfoTmp
	}
	return info
}

func extractCol(table *schemas.Table) []pojo.ColInfo {
	cols := make([]pojo.ColInfo, len(table.Columns()))
	for i, k := range table.Columns() {
		var col pojo.ColInfo
		col.Comment = k.Comment
		col.ColName = k.Name
		col.ColType = k.SQLType.Name
		col.IsPrimary = k.IsPrimaryKey
		col.Len = k.Length
		cols[i] = col
	}
	return cols
}

func openConnect(username string, password string, ipPort string, dbName string) *xorm.Engine {
	url := username + ":" + password + "@tcp(" + ipPort + ")/" + dbName + "?charset=utf8"
	engine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		panic(err)
	}
	return engine
}
