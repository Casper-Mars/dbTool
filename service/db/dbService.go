package db

import (
	"github.com/Casper-Mars/dbTool/pojo"

	"github.com/xormplus/xorm/schemas"
)

type DataBaseService interface {
	GetAllTableInfo(username string, password string, ipPort string, dbName string) ([]pojo.TableInfo, error)
	GetAllDBs(username string, password string, ipPort string) ([]string, error)
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
