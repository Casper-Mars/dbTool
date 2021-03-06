package ui

import (
	"github.com/gotk3/gotk3/gtk"
	"log"
	"strings"
)

type ExportUi struct {
	ipPort              *gtk.Entry
	username            *gtk.Entry
	password            *gtk.Entry
	dbNames             *gtk.Entry
	storeLocation       *gtk.Entry
	confirmButton       *gtk.Button
	refreshDBListButton *gtk.Button
	layout              *gtk.Grid
	dbList              *DBListWidget
	sqlType             *gtk.ComboBoxText
}

func NewExportUi() *ExportUi {
	layout, err := gtk.GridNew()
	if err != nil {
		log.Print(err)
		log.Fatal(err)
	}
	layout.SetMarginBottom(10)
	layout.SetMarginTop(10)
	layout.SetMarginStart(10)
	layout.SetMarginEnd(10)
	exportUi := ExportUi{}
	exportUi.layout = layout
	layout.SetColumnSpacing(10)
	layout.SetRowSpacing(10)
	exportUi.ipPort = exportUi.createInputFormRow("ip和端口", layout, 0)
	exportUi.username = exportUi.createInputFormRow("用户名", layout, 1)
	exportUi.password = exportUi.createInputFormRow("密码", layout, 2)
	//exportUi.storeLocation = exportUi.createInputFormRow("数据库类型", layout, 3)
	exportUi.storeLocation = exportUi.createInputFormRow("存储位置", layout, 3)

	exportUi.createSqlTypeContent(layout, 4, []string{"mysql", "mssql"})
	rb, err := gtk.ButtonNewWithLabel("刷新数据库列表")
	if err != nil {
		log.Print(err)
		log.Fatal(err)
	}
	layout.Attach(rb, 0, 5, 3, 1)
	exportUi.refreshDBListButton = rb
	dbList := NewDBListWidget()
	layout.Attach(dbList.layout, 0, 6, 3, 1)
	exportUi.dbList = dbList
	cb, err := gtk.ButtonNewWithLabel("导出")
	if err != nil {
		log.Print(err)
		log.Fatal(err)
	}
	layout.Attach(cb, 0, 7, 3, 1)
	exportUi.confirmButton = cb
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
	list := ui.dbList.getSelectedDBList()
	log.Printf("%#v", list)
	return strings.Join(list, ",")
}

func (ui ExportUi) GetStoreLocation() string {
	text, err := ui.storeLocation.GetText()
	if err != nil {
		log.Fatal(err)
	}
	return text
}

func (ui ExportUi) GetSqlType() string {
	return ui.sqlType.GetActiveText()
}

func (ui ExportUi) GetBox() *gtk.Grid {
	return ui.layout
}

func (ui ExportUi) GetConfirmButton() *gtk.Button {
	return ui.confirmButton
}

func (ui ExportUi) GetDBListRefreshButton() *gtk.Button {
	return ui.refreshDBListButton
}

func (ui ExportUi) AddDBList(dbNames []string) {
	if len(dbNames) == 0 {
		return
	}
	ui.dbList.resetDBList()
	for _, k := range dbNames {
		err := ui.dbList.AddDB(k)
		if err != nil {
			panic(err)
		}
	}
}

func (ui ExportUi) AddDBToList(dbName string) {
	err := ui.dbList.AddDB(dbName)
	if err != nil {
		log.Fatal(err)
	}
}

func (ui ExportUi) GetContent() gtk.IWidget {
	return ui.layout
}

func (ui ExportUi) createInputFormRow(labelName string, layout *gtk.Grid, rowIndex int) *gtk.Entry {
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

func (ui *ExportUi) createSqlTypeContent(layout *gtk.Grid, rowIndex int, sqlTypeList []string) {
	labelNew, err := gtk.LabelNew("数据库类型")
	if err != nil {
		log.Print(err)
		panic(err)
	}
	layout.Attach(labelNew, 0, rowIndex, 1, 1)
	textNew, err := gtk.ComboBoxTextNew()
	if err != nil {
		log.Print(err)
		panic(err)
	}
	for _, k := range sqlTypeList {
		textNew.AppendText(k)
	}
	textNew.SetActive(0)
	layout.Attach(textNew, 1, rowIndex, 1, 1)
	ui.sqlType = textNew
}
