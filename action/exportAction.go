package action

import (
	"fmt"
	"github.com/Casper-Mars/dbTool/service/db"
	"github.com/Casper-Mars/dbTool/service/export"
	"strings"
)

type ExportAction struct {
	exportService export.ExportService
	dbService     db.DataBaseService
}

func NewExportAction(exportService export.ExportService, dbService db.DataBaseService) *ExportAction {
	return &ExportAction{
		exportService: exportService,
		dbService:     dbService,
	}
}

func (receiver ExportAction) Export(ipPort string, username string, password string, dbNames string, storeLocation string) error {
	err := checkData(ipPort, username, password)
	if err != nil {
		return err
	}
	if len(dbNames) == 0 {
		return fmt.Errorf("数据库名不能为空")
	}
	dbNameArray := strings.Split(dbNames, ",")
	for _, dbName := range dbNameArray {
		tableInfos, err := receiver.dbService.GetAllTableInfo(username, password, ipPort, dbName)
		if err != nil {
			return err
		}
		receiver.exportService.Export(ipPort, username, password, dbNames, storeLocation, tableInfos)
	}
	return nil
}

func (receiver ExportAction) GetDBList(username string, password string, ipPort string) ([]string, error) {
	err := checkData(ipPort, username, password)
	if err != nil {
		return nil, err
	}
	bs, err := receiver.dbService.GetAllDBs(username, password, ipPort)
	if err != nil {
		return nil, err
	}
	return bs, nil
}

func checkData(
	ipAndPort string,
	username string,
	pwd string) error {
	if len(ipAndPort) == 0 {
		return fmt.Errorf("ip和端口不能为空")
	}
	if len(username) == 0 {
		return fmt.Errorf("用户名不能为空")
	}
	if len(pwd) == 0 {
		return fmt.Errorf("密码不能为空")
	}
	return nil
}
