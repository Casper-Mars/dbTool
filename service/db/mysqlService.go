package db

import (
	"database/sql"
	"fmt"
	"github.com/Casper-Mars/dbTool/pojo"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xormplus/xorm"
)

type MysqlService struct {
}

func (receiver MysqlService) GetAllTableInfo(username string, password string, ipPort string, dbName string) ([]pojo.TableInfo, error) {
	engine, err := xorm.NewEngine("mysql", fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8", username, password, ipPort, dbName))
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
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s)/", username, password, ipPort))
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
