package worker

import (
	"github.com/Casper-Mars/dbTool/action"
	"github.com/Casper-Mars/dbTool/service/db"
	"github.com/Casper-Mars/dbTool/service/export"
	"github.com/Casper-Mars/dbTool/ui"
	"github.com/gotk3/gotk3/gtk"
	"log"
)

type ExportWorker struct {
	AbstractWorker
}

func BuildExportWorker() Worker {
	return BuildExportWorkerWithWindow(nil)
}
func BuildExportWorkerWithWindow(app *gtk.Window) Worker {
	exportUi := ui.NewExportUi()
	exportUi.GetDBListRefreshButton().Connect("clicked", func() {
		ipPort := exportUi.GetIpPort()
		username := exportUi.GetUsername()
		pwd := exportUi.GetPassword()
		sqlType := exportUi.GetSqlType()
		log.Println("ipPort:" + ipPort)
		log.Println("username:" + username)
		log.Println("password:" + pwd)
		log.Println("sqlType:" + sqlType)
		var act *action.ExportAction
		switch sqlType {
		case "mysql":
			act = action.NewExportAction(export.WordExportService{}, db.MysqlService{})
		case "mssql":
			act = action.NewExportAction(export.WordExportService{}, db.MsService{})
		default:
			ui.ShowWarning(app, "请选择数据库类型")
			return
		}
		//list, err := act.GetDBList("root", "!Zhisheng2020", "192.168.123.155:3306")
		list, err := act.GetDBList(username, pwd, ipPort)
		if err != nil {
			if app != nil {
				ui.ShowWarning(app, err.Error())
			}
			return
		}
		if len(list) > 0 {
			for _, k := range list {
				exportUi.AddDBToList(k)
			}
		}
	})
	exportUi.GetConfirmButton().Connect("clicked", func() {
		ipPort := exportUi.GetIpPort()
		username := exportUi.GetUsername()
		password := exportUi.GetPassword()
		names := exportUi.GetDbNames()
		storeLocation := exportUi.GetStoreLocation()
		sqlType := exportUi.GetSqlType()
		log.Println("ipPort:" + ipPort)
		log.Println("username:" + username)
		log.Println("password:" + password)
		log.Println("dbNames:" + names)
		log.Println("storeLocation:" + storeLocation)
		log.Println("sqlType:" + sqlType)
		var act *action.ExportAction
		switch sqlType {
		case "mysql":
			act = action.NewExportAction(export.WordExportService{}, db.MysqlService{})
		case "mssql":
			act = action.NewExportAction(export.WordExportService{}, db.MsService{})
		default:
			ui.ShowWarning(app, "请选择数据库类型")
			return
		}
		err := act.Export(ipPort, username, password, names, storeLocation)
		if err != nil {
			if app != nil {
				ui.ShowWarning(app, err.Error())
			}
			return
		}
	})
	return ExportWorker{
		AbstractWorker{
			face: exportUi,
		},
	}
}
