package export

import "github.com/Casper-Mars/dbTool/pojo"

type ExportService interface {
	Export(dbName string, storeLocation string, tableInfos []pojo.TableInfo)
}
