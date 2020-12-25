package service

import "github.com/Casper-Mars/dbTool/pojo"

type ExportService interface {
	Export(ipPort string, username string, password string, dbName string, storeLocation string, tableInfos []pojo.TableInfo)
}
