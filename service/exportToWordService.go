package service

import (
	"github.com/Casper-Mars/dbTool/pojo"
	"log"
	"os"
	"strings"
)

type ExportToWordService struct {
}

func (export ExportToWordService) Export(ipPort string, username string, password string, dbName string, storeLocation string, tableInfos []pojo.TableInfo) {
	if checkParam(ipPort, username, password, dbName) {
		log.Println("参数不能为空")
		return
	}
	if len(storeLocation) == 0 {
		storeLocation = "." + string(os.PathSeparator)
	} else if !strings.HasSuffix(storeLocation, string(os.PathSeparator)) {
		storeLocation = storeLocation + string(os.PathSeparator)
	}
	Export(tableInfos, dbName, storeLocation)
}

func checkParam(ipPort string, username string, password string, dbNames string) bool {
	if ipPort == "" || username == "" || password == "" || dbNames == "" {
		return true
	}
	return false
}
