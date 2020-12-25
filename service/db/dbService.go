package db

import (
	"github.com/Casper-Mars/dbTool/pojo"
	_ "github.com/go-sql-driver/mysql"
)

type DataBaseService interface {
	GetAllTableInfo(username string, password string, ipPort string, dbName string) ([]pojo.TableInfo, error)
	GetAllDBs(username string, password string, ipPort string) ([]string, error)
}
