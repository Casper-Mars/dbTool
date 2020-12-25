package db

import (
	"database/sql"
	"github.com/Casper-Mars/dbTool/pojo"
	"github.com/xormplus/xorm"
	"github.com/xormplus/xorm/schemas"
)

type MysqlService struct {
}

func (receiver MysqlService) GetAllTableInfo(username string, password string, ipPort string, dbName string) ([]pojo.TableInfo, error) {
	engine, err := openConnect(username, password, ipPort, dbName)
	if err != nil {
		return nil, err
	}
	defer engine.Close()
	metas, err := engine.DBMetas()
	if err != nil {
		return nil, err
	}
	info := make([]pojo.TableInfo, len(metas))
	for i, k := range metas {
		var tableInfoTmp pojo.TableInfo
		tableInfoTmp.Comment = k.Comment
		tableInfoTmp.TableName = k.Name
		tableInfoTmp.Cols = extractCol(k)
		info[i] = tableInfoTmp
	}
	return info, nil
}

func (receiver MysqlService) GetAllDBs(username string, password string, ipPort string) ([]string, error) {
	db, err := sql.Open("mysql", username+":"+password+"@tcp("+ipPort+")/")
	if err != nil {
		return nil, err
	}
	defer db.Close()
	result, err := db.Query("show databases")
	if err != nil {
		return nil, err
	}
	dbNameList := make([]string, 0)
	var dbName string
	for result.Next() {
		err = result.Scan(&dbName)
		if err != nil {
			return nil, err
		}
		dbNameList = append(dbNameList, dbName)
	}
	return dbNameList, nil
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
func openConnect(username string, password string, ipPort string, dbName string) (*xorm.Engine, error) {
	url := username + ":" + password + "@tcp(" + ipPort + ")/" + dbName + "?charset=utf8"
	engine, err := xorm.NewEngine("mysql", url)
	if err != nil {
		return nil, err
	}
	return engine, nil
}
