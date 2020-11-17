package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
)

type ExportUi struct {
	ipPort        *gtk.Entry
	username      *gtk.Entry
	password      *gtk.Entry
	dbNames       *gtk.Entry
	storeLocation *gtk.Entry
	confirmButton *gtk.Button
	layout        *gtk.Grid
}

func NewExportUi() *ExportUi {
	layout, err := gtk.GridNew()
	if err != nil {
		log.Fatal(err)
	}
	layout.SetMarginBottom(10)
	layout.SetMarginTop(10)
	layout.SetMarginStart(10)
	layout.SetMarginEnd(10)
	//index :=0
	exportUi := ExportUi{}
	exportUi.layout = layout
	layout.SetColumnSpacing(10)
	layout.SetRowSpacing(10)
	exportUi.ipPort = createInputFormRow("ip和端口", layout, 0)
	exportUi.username = createInputFormRow("用户名", layout, 1)
	exportUi.password = createInputFormRow("密码", layout, 2)
	exportUi.dbNames = createInputFormRow("数据库s", layout, 3)
	exportUi.storeLocation = createInputFormRow("存储位置", layout, 4)
	button, err := gtk.ButtonNewWithLabel("导出")
	if err != nil {
		log.Fatal(err)
	}
	layout.Attach(button, 0, 5, 3, 1)
	exportUi.confirmButton = button
	return &exportUi
}

func (ui ExportUi) GetIpPort() string {
	text, err := ui.ipPort.GetText()
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func (ui ExportUi) GetUsername() string {
	text, err := ui.username.GetText()
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func (ui ExportUi) GetPassword() string {
	text, err := ui.password.GetText()
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func (ui ExportUi) GetDbNames() string {
	text, err := ui.dbNames.GetText()
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func (ui ExportUi) GetStoreLocation() string {
	text, err := ui.storeLocation.GetText()
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func (ui ExportUi) GetBox() *gtk.Grid {
	return ui.layout
}

func (ui ExportUi) GetConfirmButton() *gtk.Button {
	return ui.confirmButton
}

func createInputFormRow(labelName string, layout *gtk.Grid, rowIndex int) *gtk.Entry {
	labelNew, err := gtk.LabelNew(labelName)
	if err != nil {
		log.Fatal(err)
	}
	layout.Attach(labelNew, 0, rowIndex, 1, 1)
	entryNew, err := gtk.EntryNew()
	if err != nil {
		log.Fatal(err)
	}
	entryNew.SetWidthChars(100)
	layout.Attach(entryNew, 1, rowIndex, 2, 1)
	return entryNew
}
