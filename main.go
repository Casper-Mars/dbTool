package main

import (
	"github.com/Casper-Mars/dbTool/service"
	"github.com/Casper-Mars/dbTool/ui"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
)

func main() {
	/*1、初始化*/
	const appId = "org.r.tool.dbexportor"
	app, err := gtk.ApplicationNew(appId, glib.APPLICATION_FLAGS_NONE)
	if err != nil {
		log.Fatal("Could not create application.", err)
	}
	/*逻辑代码*/
	/*1.创建窗口*/
	/*2.设置属性*/
	/*3.显示窗口*/
	_, err = app.Connect("activate", func() {
		onActivate(app)
	})
	if err != nil {
		log.Fatal(err)
	}
	/*主事件循环*/
	app.Run(os.Args)
	log.Println("程序结束")
}

func onActivate(application *gtk.Application) {

	appWindow, err := gtk.ApplicationWindowNew(application)
	if err != nil {
		log.Fatal("Could not create application window.", err)
	}
	appWindow.SetTitle("数据库表导出word工具v0.1")
	appWindow.SetDefaultSize(400, 400)
	appWindow.SetPosition(gtk.WIN_POS_CENTER)
	exportUi := ui.NewExportUi()
	appWindow.Add(exportUi.GetBox())
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
		exportService := service.ExportToWordService{}
		exportService.Export(ipPort, username, password, names, storeLocation)
	})
	appWindow.ShowAll()
}
