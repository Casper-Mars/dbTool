package service

import (
	"github.com/Casper-Mars/dbTool/service/db"
	"log"
	"os"
	"strings"
)

type ExportToWordService struct {
}

func (export ExportToWordService) Export(ipPort string, username string, password string, dbNames string, storeLocation string) {

	if checkParam(ipPort, username, password, dbNames) {
		log.Println("参数不能为空")
		return
	}
	if storeLocation == "" {
		storeLocation = "./"
	} else if !strings.HasSuffix(storeLocation, string(os.PathSeparator)) {
		storeLocation = storeLocation + string(os.PathSeparator)
	}
	dbNameArray := strings.Split(dbNames, ",")
	for _, dbName := range dbNameArray {
		tableInfos := db.GetAllTableInfo(username, password, ipPort, dbName)
		Export(tableInfos, dbName, storeLocation)
	}
}

func checkParam(ipPort string, username string, password string, dbNames string) bool {
	if ipPort == "" || username == "" || password == "" || dbNames == "" {
		return true
	}
	return false
}
