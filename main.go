package main

import (
	"github.com/Casper-Mars/dbTool/service"
	"github.com/Casper-Mars/dbTool/ui"
	"github.com/gotk3/gotk3/glib"
	"github.com/gotk3/gotk3/gtk"
	"log"
	"os"
	"strings"
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
	appWindow.SetTitle("数据库表导出word工具")
	appWindow.SetDefaultSize(400, 400)

	layout, err := gtk.BoxNew(gtk.ORIENTATION_VERTICAL, 2) //以水平布局创建一个容器, 第二个参数是其中控件的像素间隔
	if err != nil {
		log.Fatal(err)
	}
	appWindow.Add(layout) //将布局添加到window中
	ipPort := ui.NewFormInputWidget("ipPort")
	layout.Add(ipPort.GetRow())
	username := ui.NewFormInputWidget("username")
	layout.Add(username.GetRow())
	password := ui.NewFormInputWidget("password")
	layout.Add(password.GetRow())
	dbNames := ui.NewFormInputWidget("数据库名称(多个用逗号隔开)")
	layout.Add(dbNames.GetRow())
	storeLocation := ui.NewFormInputWidget("存储位置")
	layout.Add(storeLocation.GetRow())

	confrimButtonRow, err := gtk.BoxNew(gtk.ORIENTATION_HORIZONTAL, 10)
	if err != nil {
		log.Fatal(err)
	}
	layout.Add(confrimButtonRow)
	confirmButton, err := gtk.ButtonNewWithLabel("导出")
	if err != nil {
		log.Fatal(err)
	}
	confrimButtonRow.Add(confirmButton)
	confirmButton.Connect("clicked", func() {
		ipPortText := ipPort.GetInputText()
		usernameText := username.GetInputText()
		passwordText := password.GetInputText()
		dbNamesText := dbNames.GetInputText()
		storeLocationText := storeLocation.GetInputText()
		log.Println("ipPort:" + ipPortText)
		log.Println("username:" + usernameText)
		log.Println("password:" + passwordText)
		log.Println("dbNames:" + dbNamesText)
		log.Println("storeLocation:" + storeLocationText)
		export(ipPortText, usernameText, passwordText, dbNamesText, storeLocationText)
	})
	appWindow.ShowAll()
}

func export(ipPort string, username string, password string, dbNames string, storeLocation string) {
	dbNameArray := strings.Split(dbNames, ",")
	for _, dbName := range dbNameArray {
		tableInfos := service.GetAllTableInfo(username, password, ipPort, dbName)
		service.Export(tableInfos, dbName, storeLocation)
	}
}
