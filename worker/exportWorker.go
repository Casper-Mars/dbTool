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
	exportAction := action.NewExportAction(export.ExportToWordService{}, db.MysqlService{})
	exportUi := ui.NewExportUi()
	exportUi.GetDBListRefreshButton().Connect("clicked", func() {
		ipPort := exportUi.GetIpPort()
		username := exportUi.GetUsername()
		pwd := exportUi.GetPassword()
		log.Println("ipPort:" + ipPort)
		log.Println("username:" + username)
		log.Println("password:" + pwd)
		//list, err := exportAction.GetDBList("root", "!Zhisheng2020", "192.168.123.155:3306")
		list, err := exportAction.GetDBList(username, pwd, ipPort)
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
		log.Println("ipPort:" + ipPort)
		log.Println("username:" + username)
		log.Println("password:" + password)
		log.Println("dbNames:" + names)
		log.Println("storeLocation:" + storeLocation)
		err := exportAction.Export(ipPort, username, password, names, storeLocation)
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
